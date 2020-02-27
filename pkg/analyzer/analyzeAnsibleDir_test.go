package analyzer

import (
	"testing"

	core "github.com/MaibornWolff/iac-count/pkg/core"
)

func TestAnalyzeAnsibleDir(t *testing.T) {
	got := analyzeAnsibleDir("test/data/")
	want := make(map[string]int)
	want[core.Roles] = 1
	want[core.Plugins] = 1

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
