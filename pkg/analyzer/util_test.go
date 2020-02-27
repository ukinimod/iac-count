package analyzer

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestNumberOfLines(t *testing.T) {
	yamlString := input.ReadFileToString("test/data/main.yml")
	got := numberOfLines(yamlString)
	want := 28

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFiletype(t *testing.T) {
	got := filetype("something/roles/a.b/tasks/main.yml")
	want := "task"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
