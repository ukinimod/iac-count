package analyzer

import (
	"testing"

	"github.com/MaibornWolff/iac-count/pkg/core"
)

func TestAnalyzeHandler(t *testing.T) {
	got := analyzeHandler("test/data/roles/example/handlers/main.yml")
	want := make(map[string]int)
	want[core.Handlers] = 2
	want[core.Complexity] = 2

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
