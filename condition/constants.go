package condition

import (
	cond "github.com/rancher/norman/condition"
)

const (
	// transition type

	Created      cond.Cond = "Created"
	RunCompleted cond.Cond = "RunCompleted"

	// done type

	Completed cond.Cond = "Completed"
	Ready     cond.Cond = "Ready"

	// error type

	Failed cond.Cond = "Failed"
)
