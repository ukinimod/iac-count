package ansible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiletype(t *testing.T) {
	got := filetype("something/roles/a.b/tasks/main.yml")
	var want NodeType = Tasks

	assert.Equal(t, want, got)
}

func TestIsRoleDir(t *testing.T) {
	got := isRoleDir("roles/a.b/defaults")
	want := true

	assert.Equal(t, want, got)
}
