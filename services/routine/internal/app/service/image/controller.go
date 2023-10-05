package image

import (
	"lifthus-auth/common/guard"
	"net/http"
	"routine/internal/aws"

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
// @Router /images/{target} [post]
// @Param target path string true "images for target"
// @Param images formData file true "images of act"
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

	target, err := aws.MapImgCategory(c.Param("target"))
	if err != nil {
		log.Printf("invalid target: %v", err)
		return c.String(http.StatusBadRequest, "invalid img target")
	}

	imgFiles := form.File["images"]
	if v := IsImageFilesValid(imgFiles); !v {
		log.Printf("invalid image files: %v", err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	rb := aws.GetRoutineBucket()
	_, locations, err := rb.UploadImagesToRoutineS3(target, imgFiles)
	if err != nil {
		log.Printf("failed to upload images to routine s3: %v", err)
		return c.String(http.StatusInternalServerError, "failed to upload images to routine s3")
	}
	return c.JSON(http.StatusCreated, locations)
}
