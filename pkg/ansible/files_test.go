package ansible

import (
	"testing"
)

func TestFiles(t *testing.T) {
	got := FilesCalculator{}.Analyze("test/data", "").Value()
	want := 12

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
