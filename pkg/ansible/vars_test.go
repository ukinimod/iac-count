package ansible

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestVarsInVars(t *testing.T) {
	path := "test/data/roles/example/vars/main.yml"
	yamlString := input.ReadFileToString(path)
	got := VarsCalculator{}.Analyze(path, yamlString).Value()
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestVarsInHostVars(t *testing.T) {
	path := "test/data/host_vars/prod.yml"
	yamlString := input.ReadFileToString(path)
	got := VarsCalculator{}.Analyze(path, yamlString).Value()
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestVarsInGroupVars(t *testing.T) {
	path := "test/data/group_vars/main.yml"
	yamlString := input.ReadFileToString(path)
	got := VarsCalculator{}.Analyze(path, yamlString).Value()
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestVarsInDefaults(t *testing.T) {
	path := "test/data/roles/example/defaults/main.yml"
	yamlString := input.ReadFileToString(path)
	got := VarsCalculator{}.Analyze(path, yamlString).Value()
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
