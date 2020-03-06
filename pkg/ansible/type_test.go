package ansible

import (
	"testing"
)

func TestFiletype(t *testing.T) {
	got := filetype("something/roles/a.b/tasks/main.yml")
	var want NodeType = Tasks

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestIsRoleDir(t *testing.T) {
	got := isRoleDir("roles/a.b/defaults")
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
