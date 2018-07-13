package mapper

import (
	"fmt"
	"sort"
	"strings"

	"regexp"

	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	"github.com/rancher/norman/types/values"
	"k8s.io/api/core/v1"
	v1types "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	exprRegexp = regexp.MustCompile("^(.*[^!])(!=|=|<|>| in | notin )(.*)$")
)

type SchedulingMapper struct {
}

func (s SchedulingMapper) FromInternal(data map[string]interface{}) {
	defer func() {
		delete(data, "nodeSelector")
		delete(data, "affinity")
	}()

	var requireAll []string
	for key, value := range convert.ToMapInterface(data["nodeSelector"]) {
		if value == "" {
			requireAll = append(requireAll, key)
		} else {
			requireAll = append(requireAll, fmt.Sprintf("%s = %s", key, value))
		}
	}

	if len(requireAll) > 0 {
		values.PutValue(data, requireAll, "scheduling", "node", "requireAll")
	}

	v, ok := data["affinity"]
	if !ok || v == nil {
		return
	}

	affinity := &v1.Affinity{}
	if err := convert.ToObj(v, affinity); err != nil {
		return
	}

	if affinity.NodeAffinity != nil {
		s.nodeAffinity(data, affinity.NodeAffinity)
	}

	if affinity.PodAffinity != nil {
		s.podAffinity(data, affinity.PodAffinity)
	}

	if affinity.PodAntiAffinity != nil {
		s.podAntiAffinity(data, affinity.PodAntiAffinity)
	}
}

func defaultTopology(key string) string {
	if key == "" {
		return "kubernetes.io/hostname"
	}
	return key
}

func (s SchedulingMapper) podAffinity(data map[string]interface{}, podAffinity *v1.PodAffinity) {
	affinity := []interface{}{}
	if podAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		for _, term := range podAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
			required := PodAffinityTermToStrings(term)
			topology := term.TopologyKey
			affinity = append(affinity, map[string]interface{}{
				"required": required,
				"topology": topology,
			})
		}
	}

	if podAffinity.PreferredDuringSchedulingIgnoredDuringExecution != nil {
		for _, preferredTerm := range podAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
			term := preferredTerm.PodAffinityTerm
			preferred := PodAffinityTermToStrings(term)
			topology := term.TopologyKey
			affinity = append(affinity, map[string]interface{}{
				"preferred": preferred,
				"topology":  topology,
			})
		}
	}
	if affinity != nil {
		values.PutValue(data, affinity, "scheduling", "pod", "affinity")
	}
}

func (s SchedulingMapper) podAntiAffinity(data map[string]interface{}, podAntiAffinity *v1.PodAntiAffinity) {
	affinity := []interface{}{}
	if podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		for _, term := range podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
			required := PodAffinityTermToStrings(term)
			topology := term.TopologyKey
			affinity = append(affinity, map[string]interface{}{
				"required": required,
				"topology": topology,
			})
		}
	}

	if podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution != nil {
		for _, preferredTerm := range podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
			term := preferredTerm.PodAffinityTerm
			preferred := PodAffinityTermToStrings(term)
			topology := term.TopologyKey
			affinity = append(affinity, map[string]interface{}{
				"preferred": preferred,
				"topology":  topology,
			})
		}
	}
	if affinity != nil {
		values.PutValue(data, affinity, "scheduling", "pod", "antiAffinity")
	}
}

func (s SchedulingMapper) nodeAffinity(data map[string]interface{}, nodeAffinity *v1.NodeAffinity) {
	var requireAll []string
	var requireAny []string
	var preferred []string

	if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		for _, term := range nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
			exprs := NodeSelectorTermToStrings(term)
			if len(exprs) == 0 {
				continue
			}
			if len(requireAny) > 0 {
				// Once any is set all new terms go to any
				requireAny = append(requireAny, strings.Join(exprs, " && "))
			} else if len(requireAll) > 0 {
				// If all is already set, we actually need to move everything to any
				requireAny = append(requireAny, strings.Join(requireAll, " && "))
				requireAny = append(requireAny, strings.Join(exprs, " && "))
				requireAll = []string{}
			} else {
				// The first term is considered all
				requireAll = exprs
			}
		}
	}

	if nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution != nil {
		sortPreferred(nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution)
		for _, term := range nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
			exprs := NodeSelectorTermToStrings(term.Preference)
			preferred = append(preferred, strings.Join(exprs, " && "))
		}
	}

	if len(requireAll) > 0 {
		values.PutValue(data, requireAll, "scheduling", "node", "requireAll")
	}
	if len(requireAny) > 0 {
		values.PutValue(data, requireAny, "scheduling", "node", "requireAny")
	}
	if len(preferred) > 0 {
		values.PutValue(data, preferred, "scheduling", "node", "preferred")
	}
}

func sortPreferred(terms []v1.PreferredSchedulingTerm) {
	sort.Slice(terms, func(i, j int) bool {
		return terms[i].Weight > terms[j].Weight
	})
}

func PodAffinityTermToStrings(term v1.PodAffinityTerm) []string {
	exprs := []string{}

	for _, expr := range term.LabelSelector.MatchExpressions {
		nextExpr := ""
		switch expr.Operator {
		case v1types.LabelSelectorOpIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s in (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s = %s", expr.Key, expr.Values[0])
			}
		case v1types.LabelSelectorOpNotIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s notin (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s != %s", expr.Key, expr.Values[0])
			}
		case v1types.LabelSelectorOpExists:
			nextExpr = expr.Key
		case v1types.LabelSelectorOpDoesNotExist:
			nextExpr = "!" + expr.Key
		}
		if nextExpr != "" {
			exprs = append(exprs, nextExpr)
		}
	}

	for k, v := range term.LabelSelector.MatchLabels {
		exprs = append(exprs, fmt.Sprintf("%s = %s", k, v))
	}
	return exprs
}

func NodeSelectorTermToStrings(term v1.NodeSelectorTerm) []string {
	exprs := []string{}

	for _, expr := range term.MatchExpressions {
		nextExpr := ""
		switch expr.Operator {
		case v1.NodeSelectorOpIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s in (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s = %s", expr.Key, expr.Values[0])
			}
		case v1.NodeSelectorOpNotIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s notin (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s != %s", expr.Key, expr.Values[0])
			}
		case v1.NodeSelectorOpExists:
			nextExpr = expr.Key
		case v1.NodeSelectorOpDoesNotExist:
			nextExpr = "!" + expr.Key
		case v1.NodeSelectorOpGt:
			if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s > %s", expr.Key, expr.Values[0])
			}
		case v1.NodeSelectorOpLt:
			if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s < %s", expr.Key, expr.Values[0])
			}
		}

		if nextExpr != "" {
			exprs = append(exprs, nextExpr)
		}
	}

	return exprs
}

func StringsToNodeSelectorTerm(exprs []string) []v1.NodeSelectorTerm {
	result := []v1.NodeSelectorTerm{}

	for _, inter := range exprs {
		term := v1.NodeSelectorTerm{}

		for _, expr := range strings.Split(inter, "&&") {
			groups := exprRegexp.FindStringSubmatch(expr)
			selectorRequirement := v1.NodeSelectorRequirement{}

			if groups == nil {
				if strings.HasPrefix(expr, "!") {
					selectorRequirement.Key = strings.TrimSpace(expr[1:])
					selectorRequirement.Operator = v1.NodeSelectorOpDoesNotExist
				} else {
					selectorRequirement.Key = strings.TrimSpace(expr)
					selectorRequirement.Operator = v1.NodeSelectorOpExists
				}
			} else {
				selectorRequirement.Key = strings.TrimSpace(groups[1])
				selectorRequirement.Values = convert.ToValuesSlice(groups[3])
				op := strings.TrimSpace(groups[2])
				switch op {
				case "=":
					selectorRequirement.Operator = v1.NodeSelectorOpIn
				case "!=":
					selectorRequirement.Operator = v1.NodeSelectorOpNotIn
				case "notin":
					selectorRequirement.Operator = v1.NodeSelectorOpNotIn
				case "in":
					selectorRequirement.Operator = v1.NodeSelectorOpIn
				case "<":
					selectorRequirement.Operator = v1.NodeSelectorOpLt
				case ">":
					selectorRequirement.Operator = v1.NodeSelectorOpGt
				}
			}

			term.MatchExpressions = append(term.MatchExpressions, selectorRequirement)
		}

		result = append(result, term)
	}

	return result
}

func StringsToLabelSelectorTerm(exprs []string) []v1types.LabelSelectorRequirement {
	result := []v1types.LabelSelectorRequirement{}

	for _, inter := range exprs {
		for _, expr := range strings.Split(inter, "&&") {
			term := v1types.LabelSelectorRequirement{}
			groups := exprRegexp.FindStringSubmatch(expr)
			if groups == nil {
				if strings.HasPrefix(expr, "!") {
					term.Key = strings.TrimSpace(expr[1:])
					term.Operator = v1types.LabelSelectorOpDoesNotExist
				} else {
					term.Key = strings.TrimSpace(expr)
					term.Operator = v1types.LabelSelectorOpExists
				}
			} else {
				term.Key = strings.TrimSpace(groups[1])
				term.Values = convert.ToValuesSlice(groups[3])
				op := strings.TrimSpace(groups[2])
				switch op {
				case "=":
					term.Operator = v1types.LabelSelectorOpIn
				case "!=":
					term.Operator = v1types.LabelSelectorOpNotIn
				case "notin":
					term.Operator = v1types.LabelSelectorOpNotIn
				case "in":
					term.Operator = v1types.LabelSelectorOpIn
				}
			}
			result = append(result, term)
		}
	}
	return result
}

func (s SchedulingMapper) ToInternal(data map[string]interface{}) {
	defer func() {
		delete(data, "scheduling")
	}()

	nodeName := convert.ToString(values.GetValueN(data, "scheduling", "node", "nodeId"))
	if nodeName != "" {
		data["nodeName"] = nodeName
	}

	requireAll := convert.ToStringSlice(values.GetValueN(data, "scheduling", "node", "requireAll"))
	requireAny := convert.ToStringSlice(values.GetValueN(data, "scheduling", "node", "requireAny"))
	preferred := convert.ToStringSlice(values.GetValueN(data, "scheduling", "node", "preferred"))

	passNodeAffinity := false
	if len(requireAll) == 0 && len(requireAny) == 0 && len(preferred) == 0 {
		passNodeAffinity = true
	}

	nodeAffinity := v1.NodeAffinity{}
	if !passNodeAffinity {
		if len(requireAll) > 0 {
			nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = &v1.NodeSelector{
				NodeSelectorTerms: []v1.NodeSelectorTerm{
					AggregateTerms(StringsToNodeSelectorTerm(requireAll)),
				},
			}
		}

		if len(requireAny) > 0 {
			if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
				nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = &v1.NodeSelector{}
			}
			nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms = StringsToNodeSelectorTerm(requireAny)
		}

		if len(preferred) > 0 {
			count := int32(100)
			for _, term := range StringsToNodeSelectorTerm(preferred) {
				nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(
					nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution, v1.PreferredSchedulingTerm{
						Weight:     count,
						Preference: term,
					})
				count--
			}
		}
	}

	podAffinity := v1.PodAffinity{}
	podAntiAffinity := v1.PodAntiAffinity{}

	pAffinity := convert.ToMapSlice(values.GetValueN(data, "scheduling", "pod", "affinity"))
	pAntiAffinity := convert.ToMapSlice(values.GetValueN(data, "scheduling", "pod", "antiAffinity"))

	passPodAffinity := false
	if (pAffinity == nil && pAntiAffinity == nil) || (len(pAffinity) == 0 && len(pAntiAffinity) == 0) {
		passPodAffinity = true
	}

	if !passPodAffinity {
		if len(pAffinity) > 0 {
			if podAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
				podAffinity.RequiredDuringSchedulingIgnoredDuringExecution = []v1.PodAffinityTerm{}
			}
			if podAffinity.PreferredDuringSchedulingIgnoredDuringExecution == nil {
				podAffinity.PreferredDuringSchedulingIgnoredDuringExecution = []v1.WeightedPodAffinityTerm{}
			}
			for _, affinity := range pAffinity {
				// hard constraints
				required := convert.ToStringSlice(values.GetValueN(convert.ToMapInterface(affinity), "required"))
				if required != nil && len(required) > 0 {
					podAffinityTerm := v1.PodAffinityTerm{
						TopologyKey: defaultTopology(""),
						LabelSelector: &v1types.LabelSelector{
							MatchExpressions: StringsToLabelSelectorTerm(required),
						},
					}
					podAffinity.RequiredDuringSchedulingIgnoredDuringExecution = append(podAffinity.RequiredDuringSchedulingIgnoredDuringExecution, podAffinityTerm)
				}

				// soft constraints
				preferred := convert.ToStringSlice(values.GetValueN(convert.ToMapInterface(affinity), "preferred"))
				if preferred != nil && len(preferred) > 0 {
					podWeightedAffinityTerm := v1.WeightedPodAffinityTerm{
						PodAffinityTerm: v1.PodAffinityTerm{
							TopologyKey: defaultTopology(""),
							LabelSelector: &v1types.LabelSelector{
								MatchExpressions: StringsToLabelSelectorTerm(preferred),
							},
						},
						Weight: int32(100),
					}
					podAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(podAffinity.PreferredDuringSchedulingIgnoredDuringExecution, podWeightedAffinityTerm)
				}
			}
		}

		if len(pAntiAffinity) > 0 {
			if podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
				podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = []v1.PodAffinityTerm{}
			}
			if podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution == nil {
				podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution = []v1.WeightedPodAffinityTerm{}
			}
			for _, affinity := range pAntiAffinity {
				// hard constraints
				required := convert.ToStringSlice(values.GetValueN(convert.ToMapInterface(affinity), "required"))
				if required != nil && len(required) > 0 {
					podAffinityTerm := v1.PodAffinityTerm{
						TopologyKey: defaultTopology(""),
						LabelSelector: &v1types.LabelSelector{
							MatchExpressions: StringsToLabelSelectorTerm(required),
						},
					}
					podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = append(podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution, podAffinityTerm)
				}

				// soft constraints
				preferred := convert.ToStringSlice(values.GetValueN(convert.ToMapInterface(affinity), "preferred"))
				podWeightedAffinityTerm := v1.WeightedPodAffinityTerm{
					PodAffinityTerm: v1.PodAffinityTerm{
						TopologyKey: defaultTopology(""),
						LabelSelector: &v1types.LabelSelector{
							MatchExpressions: StringsToLabelSelectorTerm(preferred),
						},
					},
					Weight: int32(100),
				}
				podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution, podWeightedAffinityTerm)
			}
		}
	}

	affinity, _ := convert.EncodeToMap(&v1.Affinity{
		NodeAffinity:    &nodeAffinity,
		PodAffinity:     &podAffinity,
		PodAntiAffinity: &podAntiAffinity,
	})
	if passNodeAffinity {
		values.PutValue(data, nil, "affinity", "nodeAffinity")
	} else {
		if nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution == nil {
			values.PutValue(affinity, nil, "nodeAffinity", "preferredDuringSchedulingIgnoredDuringExecution")
		}

		if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
			values.PutValue(affinity, nil, "nodeAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
		}
	}

	if passPodAffinity {
		values.PutValue(data, nil, "affinity", "podAffinity")
	} else {
		if podAffinity.PreferredDuringSchedulingIgnoredDuringExecution == nil {
			values.PutValue(affinity, nil, "podAffinity", "preferredDuringSchedulingIgnoredDuringExecution")
		}

		if podAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
			values.PutValue(affinity, nil, "podAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
		}

		if podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution == nil {
			values.PutValue(affinity, nil, "podAntiAffinity", "preferredDuringSchedulingIgnoredDuringExecution")
		}

		if podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
			values.PutValue(affinity, nil, "podAntiAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
		}
	}

	data["affinity"] = affinity
}

func AggregateTerms(terms []v1.NodeSelectorTerm) v1.NodeSelectorTerm {
	result := v1.NodeSelectorTerm{}
	for _, term := range terms {
		result.MatchExpressions = append(result.MatchExpressions, term.MatchExpressions...)
	}
	return result
}

func (s SchedulingMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	delete(schema.ResourceFields, "nodeSelector")
	delete(schema.ResourceFields, "affinity")
	return nil
}
