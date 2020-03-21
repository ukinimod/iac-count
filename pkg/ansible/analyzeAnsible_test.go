package ansible

import (
	"testing"

	"github.com/MaibornWolff/iac-count/pkg/metrics"
	"github.com/MaibornWolff/iac-count/pkg/model"
)

func TestAnalyzeAnsibleProject(t *testing.T) {
	got := AnalyzeAnsibleProject("test/data/", make([]string, 0))

	gotLen := len(got)
	wantLen := 18

	if gotLen != wantLen {
		t.Errorf("wrong number of nodes: got %d want %d", gotLen, wantLen)
	}

	gotRoot := got["."]
	gotChildrenLen := len(gotRoot.Children)
	wantChildrenLen := 4

	if gotChildrenLen != wantChildrenLen {
		t.Errorf("wrong number of children for . : got children %s wanted %d children, ", gotRoot.Children, wantChildrenLen)
	}

	gotMetrics := gotRoot.Metrics
	want := make(map[string]int)
	want["loc"] = 133
	want["rloc"] = 109
	want["comment_lines"] = 12
	want["files"] = 12

	for k := range want {
		if gotMetrics[k].Value() != want[k] {
			t.Errorf("got %d for %q in . want %d", gotMetrics[k], k, want[k])
		}
	}
}

func TestRecursiveMetricAggregation(t *testing.T) {
	given := make(map[string]model.NodeData)
	children := make(map[string]string)
	children["main.yml"] = "main.yml"
	given["."] = model.NodeData{
		Path:     ".",
		Children: children,
		Metrics:  make(map[string]metrics.Metric),
	}
	childMetrics := make(map[string]metrics.Metric)
	childMetrics["loc"] = metrics.Loc{
		Val: 1,
	}
	given["main.yml"] = model.NodeData{
		Path:    "main.yml",
		Metrics: childMetrics,
	}

	recursiveMetricAggregation(given, ".", "main.yml")

	gotMetrics := given["."].Metrics

	want := make(map[string]int)
	want["loc"] = 1

	for k := range want {
		if gotMetrics[k].Value() != want[k] {
			t.Errorf("got %d for %q in . want %d", gotMetrics[k], k, want[k])
		}
	}
}
