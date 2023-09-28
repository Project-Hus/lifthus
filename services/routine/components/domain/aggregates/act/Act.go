package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"time"
)

type ActId domain.Id
type ActCode string
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
	ImageSrcs       *ActImageSrcs
	Text            *ActText
	Characteristics *ActCharacteristics
}

func (a *Act) Upgrade(author user.User, targets ActUpgradeTargets) (*Act, error) {
	if a.Author() != author.Id() {
		return nil, domain.ErrUnauthorized
	}
	if !targets.IsValid() {
		return nil, ErrInvalidActInfo
	}

	newVer := upgradeActVersion(a, targets)

	a.versions = append(a.versions, newVer)
	return a, nil
}

func upgradeActVersion(a *Act, targets ActUpgradeTargets) *ActVersion {
	prevVer := a.LatestVersion()
	newImageSrcs := prevVer.ImageSrcs()
	newText := prevVer.Text()
	newCharacteristics := prevVer.Characteristics()
	if targets.ImageSrcs != nil {
		newImageSrcs = *targets.ImageSrcs
	}
	if targets.Text != nil {
		newText = *targets.Text
	}
	if targets.Characteristics != nil {
		newCharacteristics = *targets.Characteristics
	}
	newVer := ActVersionFrom(prevVer.Version()+1, newImageSrcs, newText, newCharacteristics, domain.CreatedAt(time.Now()))
	return newVer
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
