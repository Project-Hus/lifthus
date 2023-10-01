package dto

type CreateActDto struct {
	ActType   string   `json:"actType"`
	Name      string   `json:"name"`
	Author    string   `json:"author"`
	Text      string   `json:"text"`
	ImageSrcs []string `json:"imageSrcs,omitempty"`
}

type UpgradeActDto struct{}

type QueryActDto struct{}
