package status

import (
	"strings"

	"github.com/rancher/norman/types/convert"
	"github.com/rancher/norman/types/mapper"
)

type conditionMapping struct {
	Name        string
	State       string
	Transition  bool
	Error       bool
	FalseIsGood bool
}

type condition struct {
	Type    string
	Status  string
	Message string
}

var conditionMappings = []conditionMapping{
	{
		Name:       "Initialized",
		Transition: true,
		State:      "initializing",
	},
	{
		Name:       "Ready",
		Transition: true,
		State:      "activating",
	},
	{
		Name:       "Available",
		Transition: true,
		State:      "activating",
	},
	{
		Name:       "Progressing",
		Transition: true,
		State:      "updating",
	},
	{
		Name:       "ConfigOK",
		Transition: true,
		State:      "configuring",
	},
	{
		Name:       "PodScheduled",
		Transition: true,
		State:      "scheduling",
	},
	{
		Name:  "Completed",
		State: "completed",
	},
	{
		Name:  "Failed",
		Error: true,
		State: "error",
	},
	{
		Name:        "OutOfDisk",
		Error:       true,
		FalseIsGood: true,
	},
	{
		Name:        "MemoryPressure",
		Error:       true,
		FalseIsGood: true,
	},
	{
		Name:        "DiskPressure",
		Error:       true,
		FalseIsGood: true,
	},
	{
		Name:  "NetworkUnavailable",
		Error: true,
	},
	{
		Name:        "Unschedulable",
		Error:       true,
		FalseIsGood: true,
	},
	{
		Name:        "ReplicaFailure",
		Error:       true,
		FalseIsGood: true,
	},
}

func Set(data map[string]interface{}) {
	val, ok := mapper.GetValue(data, "status", "conditions")
	if !ok || val == nil {
		return
	}

	var conditions []condition
	if err := convert.ToObj(val, &conditions); err != nil {
		// ignore error
		return
	}

	conditionMap := map[string]condition{}
	for _, c := range conditions {
		conditionMap[c.Type] = condition{
			Status:  c.Status,
			Message: c.Message,
		}
	}

	state := ""
	error := false
	transitioning := false
	message := ""

	for _, conditionMapping := range conditionMappings {
		good := true
		condition, ok := conditionMap[conditionMapping.Name]
		if !ok {
			continue
		}

		if conditionMapping.FalseIsGood && condition.Status == "True" {
			good = false
		} else if !conditionMapping.FalseIsGood && condition.Status == "False" {
			good = false
		}

		if !good && conditionMapping.Transition {
			transitioning = true
		}

		if !good && state == "" && conditionMapping.State != "" {
			state = conditionMapping.State
		}

		if !good && conditionMapping.Error {
			error = true
			if len(message) > 0 {
				message = strings.Join([]string{message, condition.Message}, ",")
			} else {
				message = condition.Message
			}
		}
	}

	if state == "" {
		val, _ := mapper.GetValueN(data, "status", "phase").(string)
		if val != "" {
			state = val
		}
	}

	if state == "" {
		state = "active"
	}

	data["state"] = strings.ToLower(state)
	if error {
		data["transitioning"] = "error"
	} else if transitioning {
		data["transitioning"] = "yes"
	} else {
		data["transitioning"] = "no"
	}

	data["transitioningMessage"] = message
}
