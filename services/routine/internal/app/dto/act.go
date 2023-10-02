package dto

type CreateActDto struct {
	ActType   string
	Name      string
	Author    uint64
	Text      string
	ImageSrcs []string
}

type UpgradeActDto struct{}

type QueryActDto struct{}
