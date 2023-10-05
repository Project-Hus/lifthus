package image

import (
	"routine/internal/app/dto"
	"strconv"

	"github.com/labstack/echo/v4"
)

func parseFormAndGenerateCreateActDto(c echo.Context, locations []string) (*dto.CreateActDto, error) {
	authorId, err := strconv.ParseUint(c.FormValue("author"), 10, 64)
	if err != nil {
		return nil, err
	}
	return &dto.CreateActDto{
		ActType:   c.FormValue("actType"),
		Name:      c.FormValue("name"),
		Author:    authorId,
		Text:      c.FormValue("text"),
		ImageSrcs: locations,
	}, nil
}
