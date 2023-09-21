package program

import (
	"testing"
)

func TestCreateProgramFailByTitle(t *testing.T) {
	title := GetOverflownTitle()
	newProgInfo := NewProgramInfo{
		Title: title,
	}
	if IsCreationSuccess(&newProgInfo) {
		t.Errorf("expected error, but got nil")
	}
}

func TestCreateProgramFailByImageSrcs(t *testing.T) {
	imageSrcs := GetOverflownImageSrcs()
	newProgInfo := NewProgramInfo{
		ImageSrcs: imageSrcs,
	}
	if IsCreationSuccess(&newProgInfo) {
		t.Errorf("expected error, but got nil")
	}
}

func TestCreateProgramFailByDescription(t *testing.T) {
	desc := GetOverflownDescription()
	newProgInfo := NewProgramInfo{
		Description: desc,
	}
	if IsCreationSuccess(&newProgInfo) {
		t.Errorf("expected error, but got nil")
	}
}
