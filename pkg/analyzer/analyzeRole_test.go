package analyzer

import (
	"testing"

	core "github.com/ukinimod/iac-count/pkg/core"
)

func TestAnalyzeRole(t *testing.T) {
	got := analyzeRole("test/data/roles/example")
	want := make(map[string]int)
	want[core.Templates] = 1
	want[core.StaticFiles] = 1
	want[core.Plugins] = 1

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
