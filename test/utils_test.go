package test

import (
	"testing"
)

func TestNumAdd(t *testing.T) {
	got := NumAdd(1, 1)
	if got != 2 {
		t.Errorf("NumAdd(1,1) = %d, want 2", got)
	}
}
