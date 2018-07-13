package mapper

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/rancher/norman/types/values"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNodeAffinity(t *testing.T) {
	s := SchedulingMapper{}

	testInputs := []map[string]interface{}{
		{
			"scheduling": map[string]interface{}{
				"node": map[string]interface{}{
					"requireAll": []string{"foo = bar && foo1 != bar1"},
				},
			},
		},
	}

	expected := []v1.Affinity{
		{
			NodeAffinity: &v1.NodeAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: &v1.NodeSelector{
					NodeSelectorTerms: []v1.NodeSelectorTerm{
						{
							MatchExpressions: []v1.NodeSelectorRequirement{
								{
									Key:      "foo",
									Operator: v1.NodeSelectorOpIn,
									Values:   []string{"bar"},
								},
								{
									Key:      "foo1",
									Operator: v1.NodeSelectorOpNotIn,
									Values:   []string{"bar1"},
								},
							},
						},
					},
				},
			},
		},
	}
	for i, data := range testInputs {
		s.ToInternal(data)
		buf, err := json.Marshal(expected[i])
		if err != nil {
			t.Fatal(err)
		}
		var m map[string]interface{}
		if err := json.Unmarshal(buf, &m); err != nil {
			t.Fatal(err)
		}
		gotNodeAffinity, ok := values.GetValue(data, "affinity", "nodeAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
		if !ok {
			t.Fatal("key not found")
		}
		expectedNodeAffinity, ok := values.GetValue(m, "nodeAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
		if !ok {
			t.Fatal("key not found")
		}
		if !reflect.DeepEqual(gotNodeAffinity, expectedNodeAffinity) {
			t.Fatalf("test result not match! expected: %+v, got: %+v", expectedNodeAffinity, gotNodeAffinity)
		}
	}
}

func TestPodAffinity(t *testing.T) {
	s := SchedulingMapper{}

	testInputs := []map[string]interface{}{
		{
			"scheduling": map[string]interface{}{
				"pod": map[string]interface{}{
					"affinity": []map[string]interface{}{
						{
							"required": []string{"foo = bar && foo1 != bar1"},
						},
					},
				},
			},
		},
	}

	expected := []v1.Affinity{
		{
			PodAffinity: &v1.PodAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: []v1.PodAffinityTerm{
					{
						LabelSelector: &metav1.LabelSelector{
							MatchExpressions: []metav1.LabelSelectorRequirement{
								{
									Key:      "foo",
									Operator: metav1.LabelSelectorOpIn,
									Values:   []string{"bar"},
								},
								{
									Key:      "foo1",
									Operator: metav1.LabelSelectorOpNotIn,
									Values:   []string{"bar1"},
								},
							},
						},
						TopologyKey: "kubernetes.io/hostname",
					},
				},
			},
		},
	}

	for i, data := range testInputs {
		s.ToInternal(data)
		buf, err := json.Marshal(expected[i])
		if err != nil {
			t.Fatal(err)
		}
		var m map[string]interface{}
		if err := json.Unmarshal(buf, &m); err != nil {
			t.Fatal(err)
		}
		gotPodAffinity, ok := values.GetValue(data, "affinity", "podAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
		if !ok {
			t.Fatal("key not found")
		}
		expectedPodAffinity, ok := values.GetValue(m, "podAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
		if !ok {
			t.Fatal("key not found")
		}
		if !reflect.DeepEqual(gotPodAffinity, expectedPodAffinity) {
			t.Fatalf("test result not match! expected: %+v, got: %+v", expectedPodAffinity, gotPodAffinity)
		}
	}
}
