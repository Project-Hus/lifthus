package actCommand

import (
	"routine/internal/app/aws"
	"routine/internal/app/dto"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getMultipartFormAndUploadActImagesToRoutineS3(c echo.Context) (locations []string, err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}
	rb := aws.GetRoutineBucket()
	locations, err = rb.UploadActImagesToRoutineS3(form.File["images"])
	if err != nil {
		return nil, err
	}
	return locations, nil
}

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
