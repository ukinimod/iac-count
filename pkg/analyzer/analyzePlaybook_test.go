package analyzer

import (
	"testing"

	"github.com/MaibornWolff/iac-count/pkg/core"
)

func TestAnalyzePlaybook(t *testing.T) {
	got := analyzePlaybook("test/data/main.yml")
	want := make(map[string]int)
	want[core.Plays] = 1
	want[core.Complexity] = 5
	want[core.Handlers] = 1

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
