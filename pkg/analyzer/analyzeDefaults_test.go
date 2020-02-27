package analyzer

import (
	"testing"

	core "github.com/MaibornWolff/iac-count/pkg/core"
	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestAnalyzeDefaults(t *testing.T) {
	yamlString := input.ReadFileToString("test/data/roles/example/defaults/main.yml")
	got := analyzeDefaultsString(yamlString)
	want := make(map[string]int)
	want[core.Defaults] = 5
	want[core.Complexity] = 1

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
