package reader

import "testing"

func TestReadMetaFile(t *testing.T) {
	exampleData := ReadFileToString("test/data/metafile.yaml")
	got := len(ReadMetaString(exampleData).AnsibleDependencies)
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	exampleData = ReadFileToString("test/data/metafile2.yaml")
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
	exampleData := ReadFileToString("test/data/taskfile.yaml")
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
	exampleData := ReadFileToString("test/data/playbookfile.yaml")
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
	exampleData := ReadFileToString("test/data/handlerfile.yaml")
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
