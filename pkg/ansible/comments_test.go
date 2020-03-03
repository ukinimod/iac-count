package ansible

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestComments(t *testing.T) {
	path := "test/data/main.yml"
	yamlString := input.ReadFileToString(path)
	got := CommentsCalculator{}.Analyze(path, yamlString).Value()
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
