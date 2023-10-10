package image

import (
	"mime/multipart"
	"net/http"
	"strings"
)

const MAX_IMAGE_SIZE = 10 * 1024 * 1024

func IsImageFilesValid(fhs []*multipart.FileHeader) bool {
	for _, fh := range fhs {
		fb := make([]byte, MAX_IMAGE_SIZE)
		if fh.Size > MAX_IMAGE_SIZE {
			return false
		}
		f, err := fh.Open()
		if err != nil {
			return false
		}
		_, err = f.Read(fb)
		if err != nil {
			return false
		}
		if ctype := http.DetectContentType(fb); !IsTypeImage(ctype) {
			return false
		}
	}
	return true
}

func IsTypeImage(ctype string) bool {
	return strings.HasPrefix(ctype, "image")
}
