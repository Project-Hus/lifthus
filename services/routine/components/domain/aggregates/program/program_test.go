package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"testing"
	"time"
)

func TestCreateProgramFailByTitle(t *testing.T) {
	title := getOverflownTitle()
	md := MetadataFrom(nil, "abc", user.From(42), domain.TimestampsFrom(time.Now(), nil))
	cts := ContentsFrom(WeeklyType, 42, title, []string{}, "desc")
	if isCreationSuccess(md, nil, cts) {
		t.Errorf("title longer than max length should fail, but succeeded")
	}
}

func TestCreateProgramFailByImageSrcs(t *testing.T) {
	imageSrcs := getOverflownImageSrcs()
	md := MetadataFrom(nil, "abc", user.From(42), domain.TimestampsFrom(time.Now(), nil))
	cts := ContentsFrom(WeeklyType, 42, "title", imageSrcs, "desc")
	if isCreationSuccess(md, nil, cts) {
		t.Errorf("images more than max count should fail, but succeeded")
	}
}

func TestCreateProgramFailByDescription(t *testing.T) {
	desc := getOverflownDescription()
	md := MetadataFrom(nil, "abc", user.From(42), domain.TimestampsFrom(time.Now(), nil))
	cts := ContentsFrom(WeeklyType, 42, "title", []string{}, desc)
	if isCreationSuccess(md, nil, cts) {
		t.Errorf("description longer than max length should fail, but succeeded")
	}
}

func TestUpdateProgramFailByUnauthorized(t *testing.T) {
	md := MetadataFrom(nil, "abc", user.From(42), domain.TimestampsFrom(time.Now(), nil))
	cts := ContentsFrom(WeeklyType, 42, "title", []string{}, "desc")
	program := From(md, nil, cts)
	_, err := program.UpdateContents(user.From(43), ProgramContentsUpdates{})
	if err != domain.ErrUnauthorized {
		t.Errorf("update by unauthorized user should fail, but succeeded")
	}
}

func TestUpdateProgramTitleFail(t *testing.T) {}

func TestUpdateProgramDescFail(t *testing.T) {}
