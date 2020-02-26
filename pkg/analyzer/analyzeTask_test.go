package analyzer

import (
	"testing"

	"github.com/ukinimod/iac-count/pkg/core"
)

func TestAnalyzeTask(t *testing.T) {
	got := analyzeTask("test/data/roles/example/tasks/main.yml")
	want := make(map[string]int)
	want[core.Tasks] = 9
	want[core.TaggedTasks] = 3
	want[core.CustomFacts] = 0
	want[core.Complexity] = 12

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
