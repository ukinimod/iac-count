package analyzer

import (
	"testing"

	core "github.com/ukinimod/iac-count/pkg/core"
	input "github.com/ukinimod/iac-count/pkg/input"
)

func TestAnalyzeGroupVarsString(t *testing.T) {
	yamlString := input.ReadFileToString("test/data/group_vars/main.yml")
	got := analyzeGroupVarsString(yamlString)
	want := make(map[string]int)
	want[core.GroupVars] = 2
	want[core.Complexity] = 1

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
