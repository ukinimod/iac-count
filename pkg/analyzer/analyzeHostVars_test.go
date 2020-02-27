package analyzer

import (
	"testing"

	core "github.com/MaibornWolff/iac-count/pkg/core"
	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestAnalyzeHostVarsString(t *testing.T) {
	yamlString := input.ReadFileToString("test/data/host_vars/prod.yml")
	got := analyzeHostVarsString(yamlString)
	want := make(map[string]int)
	want[core.HostVars] = 3
	want[core.Complexity] = 1

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
