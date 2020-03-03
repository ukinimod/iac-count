package ansible

import (
	"testing"
)

func TestPluginsInRole(t *testing.T) {
	got := PluginsCalculator{}.Analyze("test/data/roles/example", "").Value()
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPluginsInAnsibleDir(t *testing.T) {
	got := PluginsCalculator{}.Analyze("test/data/", "").Value()
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
