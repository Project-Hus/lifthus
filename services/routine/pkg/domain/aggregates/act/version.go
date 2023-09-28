package act

import "routine/components/domain"

type ActVersionCode domain.Code
type ActVersionNumber uint
type ActImageSrcs []string
type ActText string

type ActVersion struct {
	code      ActVersionCode
	version   ActVersionNumber
	imageSrcs ActImageSrcs
	text      ActText
	createdAt domain.CreatedAt
}

func (v ActVersion) Version() ActVersionNumber {
	return v.version
}

func (v ActVersion) ImageSrcs() ActImageSrcs {
	return v.imageSrcs
}

func (v ActVersion) Text() ActText {
	return v.text
}

func (v ActVersion) CreatedAt() domain.CreatedAt {
	return v.createdAt
}
