package domain

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSuccProgramFactoryFrom(t *testing.T) {
	tmpId := uint64(42)
	expected := Program{id: &tmpId}
	result := Program{id: &tmpId}
	if diff := cmp.Diff(expected, result, cmp.AllowUnexported(Program{})); diff != "" {
		t.Error(diff)
	}
}
