package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

func init() {
	os.Setenv("AWS_ACCESS_KEY_ID", os.Getenv("LIFTHUS_ACCESS_KEY_ID"))
	os.Setenv("AWS_SECRET_ACCESS_KEY", os.Getenv("LIFTHUS_SECRET_ACCESS_KEY"))
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("couldn't load default aws configuration: %v", err)
		return
	}
	s3Client = s3.NewFromConfig(sdkConfig)
}
