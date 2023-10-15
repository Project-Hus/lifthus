package act

import (
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
	"testing"
)

func TestCreateActFailByInvalidName(t *testing.T) {
	author := user.UserFrom(42)
	_, err := CreateAct(
		WeightType,
		getTooLongActName(),
		*author,
		getValidActImages(),
		getValidActText(),
	)
	if err != ErrInvalidActInfo {
		t.Errorf("too long act name is expected to cause ErrInvalidActInfo, but got %v", err)
	}
	_, err = CreateAct(TimeType, "ts", *author, getValidActImages(), getValidActText())
	if err != ErrInvalidActInfo {
		t.Errorf("too short act name is expected to cause ErrInvalidActInfo, but got %v", err)
	}
}

func TestCreateActFailByImages(t *testing.T) {
	author := user.UserFrom(42)
	_, err := CreateAct(SimpleType,
		getValidActName(),
		*author,
		getTooManyActImages(),
		getValidActText(),
	)
	if err != ErrInvalidActInfo {
		t.Errorf("too many act images are expected to cause ErrInvalidActInfo but got %v", err)
	}
}

func TestCreateActFailByText(t *testing.T) {
	author := user.UserFrom(42)
	_, err := CreateAct(WeightType,
		getValidActName(),
		*author,
		getValidActImages(),
		getTooLongActText(),
	)
	if err != ErrInvalidActInfo {
		t.Errorf("too long act text is expected to cause ErrInvalidActInfo but got %v", err)
	}
	_, err = CreateAct(WeightType, getValidActName(), *author, getValidActImages(), "ts")
	if err != ErrInvalidActInfo {
		t.Errorf("too short act text is expected to cause ErrInvalidActInfo but got %v", err)
	}
}

func TestUpdateActSucc(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	newText := ActText(getValidActText() + "Hello!")
	newImageSrcs := ActImageSrcs{"https://example.com/updated_image.png"}
	updates := ActUpdateTargets{ImageSrcs: &newImageSrcs, Text: &newText}
	updatedActRef, err := act.Update(*author, updates)
	if err != nil {
		t.Errorf("updating act failed with error:%v", err)
	}
	if act.Text() != newText || updatedActRef.Text() != newText {
		t.Errorf("act text is not updated")
	}
	if act.ImageSrcs()[0] != newImageSrcs[0] || updatedActRef.ImageSrcs()[0] != newImageSrcs[0] {
		t.Errorf("act images are not updated")
	}
}

func TestUpdateActFailByUnauthorized(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	abuser := user.UserFrom(43)
	newText := ActText(getValidActText() + "Hello!")
	updates := ActUpdateTargets{Text: &newText}
	_, err := act.Update(*abuser, updates)
	if err != domain.ErrUnauthorized {
		t.Errorf("update by unauthorized user is expected to cause ErrUnauthorized but got %v", err)
	}
	if act.Text() == newText {
		t.Errorf("act text should not be updated")
	}
}

func TestUpdateActFailByImages(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	newImageSrcs := getTooManyActImages()
	updates := ActUpdateTargets{ImageSrcs: &newImageSrcs}
	_, err := act.Update(*author, updates)
	if err != ErrInvalidActInfo {
		t.Errorf("update with too many images is expected to cause ErrInvalidActInfo but got %v", err)
	}
	if len(act.ImageSrcs()) == len(newImageSrcs) {
		t.Errorf("act images should not be updated")
	}
}

func TestUpdateActFailByText(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	newText := ActText(getTooLongActText())
	updates := ActUpdateTargets{Text: &newText}
	_, err := act.Update(*author, updates)
	if err != ErrInvalidActInfo {
		t.Errorf("update with too long text is expected to cause ErrInvalidActInfo but got %v", err)
	}
	if act.Text() == newText {
		t.Errorf("act text should not be updated")
	}
}
