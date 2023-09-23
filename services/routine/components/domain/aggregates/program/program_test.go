package program

import (
	"routine/components/domain/aggregates/user"
	"testing"
)

func TestCreateProgramFailByTitle(t *testing.T) {
	ds := DescriptionsFrom(getOverflownTitle(), getNormalImageSrcs(), getNormalDescription())
	if isCreationSuccess(user.From(42), nil, TypeFrom(WeeklyType, 10), ds) {
		t.Errorf("title longer than max length should fail, but succeeded")
	}
}

func TestCreateProgramFailByImageSrcs(t *testing.T) {
	ds := DescriptionsFrom(getNormalTitle(), getOverflownImageSrcs(), getNormalDescription())
	if isCreationSuccess(user.From(42), nil, TypeFrom(WeeklyType, 10), ds) {
		t.Errorf("images more than max count should fail, but succeeded")
	}
}

func TestCreateProgramFailByDescription(t *testing.T) {
	ds := DescriptionsFrom(getNormalTitle(), getNormalImageSrcs(), getOverflownDescription())
	if isCreationSuccess(user.From(42), nil, TypeFrom(WeeklyType, 10), ds) {
		t.Errorf("description longer than max length should fail, but succeeded")
	}
}
