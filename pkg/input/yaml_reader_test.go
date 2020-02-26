package reader

import "testing"

func TestReadYamlAsList(t *testing.T) {
	exampleData := ReadFileToString("test/data/taskfile.yaml")
	got := len(ReadYamlAsList(exampleData))
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestReadYamlAsListOnFailure(t *testing.T) {
	invalidYaml := " as as as "
	got := len(ReadYamlAsList(invalidYaml))
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestReadYamlAsMap(t *testing.T) {
	exampleData := ReadFileToString("test/data/varsfile.yaml")
	got := ReadYamlAsMap(exampleData)["bla"]
	want := "blabla"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestReadYamlAsMapOnFailure(t *testing.T) {
	invalidYaml := " as as as"
	got := len(ReadYamlAsMap(invalidYaml))
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
