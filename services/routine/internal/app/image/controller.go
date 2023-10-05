package image

import (
	"lifthus-auth/common/guard"
	"net/http"
	"routine/internal/app/aws"

	"log"

	"github.com/labstack/echo/v4"
)

func SetImageControllerTo(e *echo.Echo) *echo.Echo {
	ac := &imageController{svc: newImageService()}
	e.POST("/routine/images/:for", ac.uploadImages, guard.UserGuard)
	return e
}

type imageController struct {
	svc *imageService
}

// uploadImages godoc
// @Router /images/{for} [post]
// @Param for path string true "images for"
// @Param images formData file false "images of act"
// @Param Authorization header string true "lifthus_st"
// @Summary
// @Tags
// Success 201 "returns locations of images"
// Failure 400 "invalid request"
// Failure 401 "unauthorized"
// Failure 500 "failed to upload images to routine s3"
func (ac *imageController) uploadImages(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("failed to parse multipart form: %v", err)
		return c.String(http.StatusBadRequest, "invalid request")
	}
	rb := aws.GetRoutineBucket()
	locations, err := rb.UploadActImagesToRoutineS3(form.File["images"])
	if err != nil {
		log.Printf("failed to upload images to routine s3: %v", err)
		return c.String(http.StatusInternalServerError, "failed to upload images to routine s3")
	}
	return c.JSON(http.StatusCreated, locations)
}
