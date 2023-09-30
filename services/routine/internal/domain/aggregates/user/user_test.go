package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSuccUserFactoryFrom(t *testing.T) {
	expected := User{42}
	result := UserFrom(42)
	if diff := cmp.Diff(expected, *result, cmp.AllowUnexported(User{})); diff != "" {
		t.Error(diff)
	}
}

func TestFailUserFactoryFrom(t *testing.T) {
	expected := User{42}
	result := UserFrom(43)
	if diff := cmp.Diff(expected, *result, cmp.AllowUnexported(User{})); diff == "" {
		t.Error(diff)
	}
}
