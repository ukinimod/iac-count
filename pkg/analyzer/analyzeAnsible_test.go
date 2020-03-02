package analyzer

import (
	"testing"

	core "github.com/MaibornWolff/iac-count/pkg/core"
)

func TestAnalyzeAnsibleProject(t *testing.T) {
	got := AnalyzeAnsibleProject("test/data/")

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
	want[core.Loc] = 138
	want[core.Tasks] = 12
	want[core.Handlers] = 3
	want[core.Complexity] = 24
	want[core.Vars] = 5
	want[core.Defaults] = 5
	want[core.GroupVars] = 2
	want[core.HostVars] = 3
	want[core.Dependencies] = 3
	want[core.Rloc] = 109
	want[core.CommentLines] = 12
	want[core.Files] = 12
	want[core.Templates] = 1
	want[core.StaticFiles] = 1
	want[core.Plugins] = 2
	want[core.Roles] = 1
	want[core.Plays] = 1
	want[core.CustomFacts] = 0
	want[core.TaggedTasks] = 3

	for k := range want {
		if gotMetrics[k] != want[k] {
			t.Errorf("got %d for %q in . want %d", gotMetrics[k], k, want[k])
		}
	}
}

func TestRecursiveMetricAggregation(t *testing.T) {
	given := make(map[string]core.FileData)
	children := make(map[string]string)
	children["main.yml"] = "main.yml"
	given["."] = core.FileData{
		RelativePath: ".",
		Children:     children,
		Metrics:      make(map[string]int),
	}
	childMetrics := make(map[string]int)
	childMetrics[core.Loc] = 1
	given["main.yml"] = core.FileData{
		RelativePath: "main.yml",
		Metrics:      childMetrics,
	}

	recursiveMetricAggregation(given, ".", "main.yml")

	gotMetrics := given["."].Metrics

	want := make(map[string]int)
	want[core.Loc] = 1

	for k := range want {
		if gotMetrics[k] != want[k] {
			t.Errorf("got %d for %q in . want %d", gotMetrics[k], k, want[k])
		}
	}
}
