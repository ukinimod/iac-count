package analyzer

import (
	"testing"

	input "github.com/ukinimod/iac-count/pkg/input"
)

func TestNumberOfLines(t *testing.T) {
	yamlString := input.ReadFileToString("test/data/main.yml")
	got := numberOfLines(yamlString)
	want := 28

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
