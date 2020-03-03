package ansible

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestHandlers(t *testing.T) {
	path := "test/data/roles/example/handlers/main.yml"
	yamlString := input.ReadFileToString(path)
	got := HandlersCalculator{}.Analyze(path, yamlString).Value()
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
