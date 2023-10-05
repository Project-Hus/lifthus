package aws

import (
	"context"
	"mime/multipart"
	"routine/internal/domain"
	"routine/pkg/helper"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetRoutineBucket() *RoutineBucket {
	return &RoutineBucket{S3Client: s3Client}
}

type RoutineBucket struct {
	S3Client *s3.Client
}

func (rb RoutineBucket) UploadImagesToRoutineS3(target ImgCategory, imgHds []*multipart.FileHeader) (keys []string, locations []string, err error) {
	return rb.UploadMultipartFilesToRoutineS3(target.Category(), imgHds)
}

// UploadMultipartFilesToRoutineS3 uploads multipart files and returns the locations of the files
func (rb RoutineBucket) UploadMultipartFilesToRoutineS3(category string, files []*multipart.FileHeader) (keys []string, locations []string, err error) {
	keys = make([]string, len(files))
	locations = make([]string, len(files))

	wg := helper.WaitGroupWaiting(len(files))
	errChan := make(chan error)

	for i, file := range files {
		go func(i int, file *multipart.FileHeader) {
			defer wg.Done()
			key, location, err := rb.UploadMultipartFileToRoutineS3(category, file)
			if err != nil {
				select {
				case errChan <- err:
					close(errChan)
				default:
				}
				return
			}
			keys[i] = key
			locations[i] = location
		}(i, file)
	}

	wg.Wait()
	select {
	case err := <-errChan:
		// TODO: delete uploaded files
		return nil, nil, err
	default:
		return keys, locations, nil
	}
}

func (rb RoutineBucket) UploadMultipartFileToRoutineS3(category string, mfh *multipart.FileHeader) (key string, location string, err error) {
	key, err = rb.generateObjKeyForFilename(category, mfh.Filename)
	if err != nil {
		return "", "", err
	}
	file, err := mfh.Open()
	if err != nil {
		return "", "", err
	}
	defer file.Close()
	_, err = rb.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(ROUTINE_BUCKET_NAME),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", "", err
	}
	location = ROUTINE_BUCKET_URL + key
	return key, location, nil
}

func (rb RoutineBucket) generateObjKeyForFilename(category string, filename string) (string, error) {
	return generateObjKeyForFilename("routine/image/"+category+"/", filename)
}

func generateObjKeyForFilename(basekey string, fn string) (okey string, err error) {
	code, err := domain.RandomHex(4)
	if err != nil {
		return "", err
	}
	basekey = helper.TrimSlash(basekey)
	basekey += "/"
	okey = basekey + strconv.FormatInt(time.Now().Unix(), 10) + "_" + code + "_" + fn
	return okey, nil
}
