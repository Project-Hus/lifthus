package aws

import (
	"context"
	"fmt"
	"mime/multipart"
	"routine/internal/domain"
	"strconv"
	"sync"
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

// UploadFile reads from a file and puts the data into an object in a bucket.
func (rb RoutineBucket) UploadMultipartFilesToRoutineS3(files []*multipart.FileHeader) ([]string, error) {
	locations := make([]string, len(files))
	wg := &sync.WaitGroup{}
	wg.Add(len(files))
	for i, file := range files {
		go func(i int, file *multipart.FileHeader) {
			defer wg.Done()
			location, err := rb.UploadMultipartFileToRoutineS3(file)
			if err != nil {
				return
			}
			locations[i] = location
		}(i, file)
	}
	wg.Wait()
	for _, location := range locations {
		if location == "" {
			return nil, fmt.Errorf("failed to upload file to s3")
		}
	}
	return locations, nil
}

func (rb RoutineBucket) UploadMultipartFileToRoutineS3(mfh *multipart.FileHeader) (string, error) {
	okey, err := rb.generateObjKeyForFilename("routine/images/act/", mfh.Filename)
	if err != nil {
		return "", err
	}
	location := "https://lifthus-routine-bucket.s3.us-west-2.amazonaws.com/" + okey
	file, err := mfh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = rb.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("lifthus-routine-bucket"),
		Key:    &okey,
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	// get the location of the s3 object
	return location, nil
}

func (rb RoutineBucket) generateObjKeyForFilename(basekey string, fn string) (string, error) {
	code, err := domain.RandomHex(4)
	if err != nil {
		return "", err
	}
	if basekey[len(basekey)-1] != '/' {
		basekey = basekey + "/"
	}
	okey := basekey + strconv.FormatInt(time.Now().Unix(), 10) + "_" + code + "_" + fn
	return okey, nil
}
