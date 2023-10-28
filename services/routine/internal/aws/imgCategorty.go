package aws

import (
	"fmt"
)

func MapImgCategory(target string) (ImgCategory, error) {
	switch target {
	case "act":
		return ImgForAct, nil
	case "program":
		return ImgForProgram, nil
	default:
		return ImgCategory{}, fmt.Errorf("invalid img target")
	}
}

type ImgCategory struct {
	category string
}

func (it ImgCategory) Category() string {
	return it.category
}

var ImgForAct = ImgCategory{category: "act"}
var ImgForProgram = ImgCategory{category: "program"}
