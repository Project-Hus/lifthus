package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"testing"
)

func TestCreateActFailByInvalidName(t *testing.T) {
	author := user.UserFrom(42)
	desc := getValidActDescription()
	_, err := CreateWeightAct(getTooLongActName(), *author, desc)
	if err != ErrInvalidActInfo {
		t.Errorf("too long act name is expected to cause ErrInvalidActInfo, but got %v", err)
	}
	desc = getValidActDescription()
	_, err = CreateWeightAct("ts", *author, desc)
	if err != ErrInvalidActInfo {
		t.Errorf("too short act name is expected to cause ErrInvalidActInfo, but got %v", err)
	}
}

func TestCreateActFailByImages(t *testing.T) {
	author := user.UserFrom(42)
	desc := getValidActDescription()
	desc.ImageSrcs = getTooManyActImages()
	_, err := CreateWeightAct(getValidActName(), *author, desc)
	if err != ErrInvalidActInfo {
		t.Errorf("too many act images are expected to cause ErrInvalidActInfo but got %v", err)
	}
}

func TestCreateActFailByText(t *testing.T) {
	author := user.UserFrom(42)
	desc := getValidActDescription()
	desc.Text = getTooLongActText()
	_, err := CreateWeightAct(getValidActName(), *author, desc)
	if err != ErrInvalidActInfo {
		t.Errorf("too long act text is expected to cause ErrInvalidActInfo but got %v", err)
	}
	desc = getValidActDescription()
	desc.Text = "ts"
	_, err = CreateWeightAct(getValidActName(), *author, desc)
	if err != ErrInvalidActInfo {
		t.Errorf("too short act text is expected to cause ErrInvalidActInfo but got %v", err)
	}
}

func TestUpdateActSucc(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	originalActVer := act.Base().Version
	newText := ActText(getValidActText() + "Hello!")
	newImageSrcs := []ActImageSrc{ActImageSrc("https://example.com/updated_image.png")}
	updates := ActUpdates{ImageSrcs: &newImageSrcs, Text: &newText}
	updatedActRef, err := act.Update(*author, updates)
	if err != nil {
		t.Errorf("updating act failed with error:%v", err)
	}
	if act.Description().Text != newText || updatedActRef.Description().Text != newText {
		t.Errorf("act text is not updated")
	}
	if act.Description().ImageSrcs[0] != newImageSrcs[0] || updatedActRef.Description().ImageSrcs[0] != newImageSrcs[0] {
		t.Errorf("act images are not updated")
	}
	if act.UpdateTargets().Version == originalActVer {
		t.Errorf("act version is not updated")
	}
}

func TestUpdateActFailByUnauthorized(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	abuser := user.UserFrom(43)
	newText := ActText(getValidActText())
	updates := ActUpdates{Text: &newText}
	_, err := act.Update(*abuser, updates)
	if err != domain.ErrUnauthorized {
		t.Errorf("update by unauthorized user is expected to cause ErrUnauthorized but got %v", err)
	}
}

func TestUpdateActFailByImages(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	newText := ActText(getValidActText() + "Hello!")
	newImageSrcs := getTooManyActImages()
	updates := ActUpdates{ImageSrcs: &newImageSrcs, Text: &newText}
	_, err := act.Update(*author, updates)
	if err != ErrInvalidActInfo {
		t.Errorf("updating act failed with error:%v", err)
	}
	if len(act.Description().ImageSrcs) == len(newImageSrcs) {
		t.Errorf("act images should not be updated")
	}
}

func TestUpdateActFailByText(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	newText := ActText(getTooLongActText())
	updates := ActUpdates{Text: &newText}
	_, err := act.Update(*author, updates)
	if err != ErrInvalidActInfo {
		t.Errorf("updating act failed with error:%v", err)
	}
	if act.Description().Text == newText {
		t.Errorf("act text should not be updated")
	}
}

func TestDeleteActFailByUnauthorized(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	abuser := user.UserFrom(43)
	_, err := act.Delete(*abuser)
	if err != domain.ErrUnauthorized {
		t.Errorf("delete by unauthorized user is expected to cause ErrUnauthorized but got %v", err)
	}
}
