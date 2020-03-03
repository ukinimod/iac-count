package ansible

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestReadMetaFile(t *testing.T) {
	exampleData := input.ReadFileToString("nonexistent")
	got := len(ReadMetaString(exampleData).AnsibleDependencies)
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	exampleData = input.ReadFileToString("test/data/roles/example/meta/main.yml")
	got = len(ReadMetaString(exampleData).AnsibleDependencies)
	want = 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	invalidYaml := " bla bla bla "
	got = len(ReadMetaString(invalidYaml).AnsibleDependencies)
	want = 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestReadTaskFile(t *testing.T) {
	exampleData := input.ReadFileToString("test/data/roles/example/tasks/main.yml")
	got := len(ReadTasksString(exampleData))
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	invalidYaml := " bla bla bla "
	got = len(ReadTasksString(invalidYaml))
	want = 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestReadPlayFile(t *testing.T) {
	exampleData := input.ReadFileToString("test/data/main.yml")
	got := len(ReadPlaybookString(exampleData))
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	invalidYaml := " bla bla bla "
	got = len(ReadPlaybookString(invalidYaml))
	want = 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestReadHandlerFile(t *testing.T) {
	exampleData := input.ReadFileToString("test/data/roles/example/handlers/main.yml")
	got := len(ReadHandlersString(exampleData))
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	invalidYaml := " bla bla bla "
	got = len(ReadHandlersString(invalidYaml))
	want = 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
