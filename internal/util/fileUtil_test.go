package util

import (
	"testing"
)

func TestRecursiveFileCount(t *testing.T) {
	got := RecursiveFileCount("nonexistingfile")
	want := 0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = RecursiveFileCount("fileUtil.go")
	want = 1

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
