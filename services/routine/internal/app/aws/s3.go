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

func (rb RoutineBucket) UploadActImagesToRoutineS3(files []*multipart.FileHeader) (locations []string, err error) {
	return rb.UploadMultipartFilesToRoutineS3("act", files)
}

// UploadMultipartFilesToRoutineS3 uploads multipart files and returns the locations of the files
func (rb RoutineBucket) UploadMultipartFilesToRoutineS3(category string, files []*multipart.FileHeader) (locations []string, err error) {
	locations = make([]string, len(files))

	wg := helper.WaitGroupWaiting(len(files))
	errChan := make(chan error)
	for i, file := range files {
		go func(i int, file *multipart.FileHeader) {
			defer wg.Done()
			location, err := rb.UploadMultipartFileToRoutineS3(category, file)
			if err != nil {
				select {
				case errChan <- err:
					close(errChan)
				default:
				}
				return
			}
			locations[i] = location
		}(i, file)
	}
	wg.Wait()
	select {
	case err := <-errChan:
		// TODO: delete uploaded files
		return nil, err
	default:
		return locations, nil
	}
}

func (rb RoutineBucket) UploadMultipartFileToRoutineS3(category string, mfh *multipart.FileHeader) (location string, err error) {
	okey, err := rb.generateObjKeyForFilename(category, mfh.Filename)
	if err != nil {
		return "", err
	}
	file, err := mfh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = rb.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(ROUTINE_BUCKET_NAME),
		Key:    aws.String(okey),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	location = ROUTINE_BUCKET_URL + okey
	return location, nil
}

func (rb RoutineBucket) generateObjKeyForFilename(category string, filename string) (string, error) {
	return generateObjKeyForFilename("routine/images/"+category+"/", filename)
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
