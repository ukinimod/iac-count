package analyzer

import (
	"testing"

	core "github.com/MaibornWolff/iac-count/pkg/core"
)

func TestAnalyzeDir(t *testing.T) {
	got := analyzeDir("test/data")
	want := make(map[string]int)
	want[core.Files] = 12

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
