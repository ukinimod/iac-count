package reader

import "testing"

func TestReadFileToString(t *testing.T) {
	exampleData := ReadFileToString("test/data/taskfile.yaml")
	got := len(exampleData)
	want := 852

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestReadFileToStringUnknownFile(t *testing.T) {
	exampleData := ReadFileToString("test/data/unknownfile")
	got := exampleData
	want := ""

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
