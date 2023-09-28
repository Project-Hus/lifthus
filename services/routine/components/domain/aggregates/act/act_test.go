package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"testing"
	"time"
)

func TestCreateActFailByInvalidName(t *testing.T) {
	author := user.UserFrom(42)
	_, err := CreateAct(WeightType, getTooLongActName(), *author, getValidActImages(), getValidActText(), getValidCharacteristics())
	if err != ErrInvalidActInfo {
		t.Errorf("too long act name is expected to cause ErrInvalidActInfo, but got %v", err)
	}
	_, err = CreateAct(TimeType, "ts", *author, getValidActImages(), getValidActText(), getValidCharacteristics())
	if err != ErrInvalidActInfo {
		t.Errorf("too short act name is expected to cause ErrInvalidActInfo, but got %v", err)
	}
}

func TestCreateActFailByImages(t *testing.T) {
	author := user.UserFrom(42)
	_, err := CreateAct(SimpleType, getValidActName(), *author, getTooManyActImages(), getValidActText(), getValidCharacteristics())
	if err != ErrInvalidActInfo {
		t.Errorf("too many act images are expected to cause ErrInvalidActInfo but got %v", err)
	}
}

func TestCreateActFailByText(t *testing.T) {
	author := user.UserFrom(42)
	_, err := CreateAct(WeightType, getValidActName(), *author, getValidActImages(), getTooLongActText(), getValidCharacteristics())
	if err != ErrInvalidActInfo {
		t.Errorf("too long act text is expected to cause ErrInvalidActInfo but got %v", err)
	}
	_, err = CreateAct(WeightType, getValidActName(), *author, getValidActImages(), "ts", getValidCharacteristics())
	if err != ErrInvalidActInfo {
		t.Errorf("too short act text is expected to cause ErrInvalidActInfo but got %v", err)
	}
}

func TestUpgradeActSucc(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	originalLatestVerNum := act.LatestVersion().Version()
	newText := ActText(getValidActText() + "Hello!")
	newImageSrcs := ActImageSrcs{"https://example.com/updated_image.png"}
	upgrades := ActUpgradeTargets{ImageSrcs: &newImageSrcs, Text: &newText}
	upgradedActRef, err := act.Upgrade(*author, upgrades)
	if err != nil {
		t.Errorf("updating act failed with error:%v", err)
	}
	if act.LatestVersion().Text() != newText || upgradedActRef.LatestVersion().Text() != newText {
		t.Errorf("act text is not updated")
	}
	if act.LatestVersion().ImageSrcs()[0] != newImageSrcs[0] || upgradedActRef.LatestVersion().ImageSrcs()[0] != newImageSrcs[0] {
		t.Errorf("act images are not updated")
	}
	if act.LatestVersion().Version() != originalLatestVerNum+1 {
		t.Errorf("act version is not updated")
	}
}

func TestUpgradeActFailByUnauthorized(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	abuser := user.UserFrom(43)
	newText := ActText(getValidActText() + "Hello!")
	upgrades := ActUpgradeTargets{Text: &newText}
	_, err := act.Upgrade(*abuser, upgrades)
	if err != domain.ErrUnauthorized {
		t.Errorf("upgrade by unauthorized user is expected to cause ErrUnauthorized but got %v", err)
	}
	if act.LatestVersion().Text() == newText {
		t.Errorf("act text should not be updated")
	}
}

func TestUpgradeActFailByImages(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	newImageSrcs := getTooManyActImages()
	upgrades := ActUpgradeTargets{ImageSrcs: &newImageSrcs}
	_, err := act.Upgrade(*author, upgrades)
	if err != ErrInvalidActInfo {
		t.Errorf("update with too many images is expected to cause ErrInvalidActInfo but got %v", err)
	}
	if len(act.LatestVersion().ImageSrcs()) == len(newImageSrcs) {
		t.Errorf("act images should not be updated")
	}
}

func TestUpgradeActFailByText(t *testing.T) {
	author := user.UserFrom(42)
	act := getValidActWithAuthor(*author)
	newText := ActText(getTooLongActText())
	upgrades := ActUpgradeTargets{Text: &newText}
	_, err := act.Upgrade(*author, upgrades)
	if err != ErrInvalidActInfo {
		t.Errorf("update with too long text is expected to cause ErrInvalidActInfo but got %v", err)
	}
	if act.LatestVersion().Text() == newText {
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

func TestActFrom(t *testing.T) {
	v1 := &ActVersion{version: 1}
	v3 := &ActVersion{version: 3}
	v4 := &ActVersion{version: 4}
	v6 := &ActVersion{version: 6}
	v11 := &ActVersion{version: 11}
	versions := ActVersions{v1, v3, v4, v6, v11}
	_, err := ActFrom("code", SimpleType, getValidActName(), *user.UserFrom(42), domain.CreatedAt(time.Now()), versions)
	if err != nil {
		t.Errorf("sorted versions should be valid but got unexpected error:%v", err)
	}
	versions = ActVersions{v1, v4, v3, v6, v11}
	_, err = ActFrom("code", SimpleType, getValidActName(), *user.UserFrom(42), domain.CreatedAt(time.Now()), versions)
	if err != ErrInvalidActVersions {
		t.Errorf("unsorted versions should cause ErrInvalidActVersions but got %v", err)
	}
	versions = ActVersions{v1, v3, v3, v6, v11}
	_, err = ActFrom("code", SimpleType, getValidActName(), *user.UserFrom(42), domain.CreatedAt(time.Now()), versions)
	if err != ErrInvalidActVersions {
		t.Errorf("versions with duplicated version number should cause ErrInvalidActVersions but got %v", err)
	}
}
