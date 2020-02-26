package analyzer

import (
	"testing"

	core "github.com/ukinimod/iac-count/pkg/core"
	input "github.com/ukinimod/iac-count/pkg/input"
)

func TestAnalyzeYamlString(t *testing.T) {
	yamlString := input.ReadFileToString("test/data/main.yml")
	got := analyzeYamlString(yamlString)
	want := make(map[string]int)
	want[core.Rloc] = 26
	want[core.CommentLines] = 2
	want[core.MaxNL] = 6
	want[core.Loc] = 28
	want[core.Complexity] = 1
	want[core.Files] = 1

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
