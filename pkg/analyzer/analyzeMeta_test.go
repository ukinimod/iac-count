package analyzer

import (
	"testing"

	"github.com/MaibornWolff/iac-count/pkg/core"
)

func TestAnalyzeMeta(t *testing.T) {
	got := analyzeMeta("test/data/roles/example/meta/main.yml")
	want := make(map[string]int)
	want[core.Dependencies] = 3
	want[core.Complexity] = 1

	for k := range want {
		if got[k] != want[k] {
			t.Errorf("got %d for %q want %d", got[k], k, want[k])
		}
	}
}
