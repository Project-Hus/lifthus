package act

import (
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
	"time"
)

type ActCode domain.Code
type ActName string
type ActVersions []*ActVersion

type Act struct {
	code ActCode

	actType   ActType
	name      ActName
	author    user.UserId
	createdAt domain.CreatedAt

	versions ActVersions
}

type ActUpgradeTargets struct {
	ImageSrcs *ActImageSrcs
	Text      *ActText
}

func (a *Act) Upgrade(author user.User, targets ActUpgradeTargets) (*Act, error) {
	if a.Author() != author.Id() {
		return nil, domain.ErrUnauthorized
	}
	if !targets.IsValid() {
		return nil, ErrInvalidActInfo
	}

	newVer, err := upgradeActVersion(a, targets)
	if err != nil {
		return nil, err
	}

	a.versions = append(a.versions, newVer)
	return a, nil
}

func upgradeActVersion(a *Act, targets ActUpgradeTargets) (*ActVersion, error) {
	prevVer := a.LatestVersion()
	newImageSrcs := prevVer.ImageSrcs()
	newText := prevVer.Text()
	if targets.ImageSrcs != nil {
		newImageSrcs = *targets.ImageSrcs
	}
	if targets.Text != nil {
		newText = *targets.Text
	}

	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}

	newVer := ActVersionFrom(ActVersionCode(code), prevVer.Version()+1, newImageSrcs, newText, domain.CreatedAt(time.Now()))
	return newVer, nil
}

func (a *Act) Delete(deleter user.User) (*Act, error) {
	if a.Author() != deleter.Id() {
		return nil, domain.ErrUnauthorized
	}
	return a, nil
}

func (a *Act) LatestVersion() *ActVersion {
	versions := a.Versions()
	vlen := len(versions)
	if vlen == 0 {
		return nil
	}
	return versions[vlen-1]
}

func (a Act) Code() ActCode {
	return a.code
}

func (a Act) Type() ActType {
	return a.actType
}

func (a Act) Name() ActName {
	return a.name
}

func (a Act) Author() user.UserId {
	return a.author
}

func (a Act) CreatedAt() domain.CreatedAt {
	return a.createdAt
}

func (a Act) Versions() []*ActVersion {
	return a.versions
}
