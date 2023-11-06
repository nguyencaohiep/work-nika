package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	_accessKeyId     = "AKIAQETG7W4XMAXEPWYV"
	_secretAccessKey = "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"
	_s3Region        = "ap-northeast-1"
	_s3Bucket        = "gear5"
)

func configS3() *s3.Client {
	creds := credentials.NewStaticCredentialsProvider(_accessKeyId, _secretAccessKey, "")
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(_s3Region),
	)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(" S3 has been initializd ! ")
	}

	return s3.NewFromConfig(cfg)
}

func UploadImageToS3(uploader *manager.Uploader, key, imageURL string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	client := &http.Client{Transport: tr}
	request, err := http.NewRequest("GET", imageURL, nil)
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println(response.StatusCode)
	}

	uploadResult, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(_s3Bucket),
		Key:         aws.String(key),
		Body:        response.Body,
		ContentType: aws.String("image/png"),
	})
	if err != nil {
		fmt.Printf("Error: %v  \n", err)
		return err
	}
	fmt.Println("=>>", uploadResult.Location)

	return nil
}

func main() {
	errorList := []string{}
	uploader := manager.NewUploader(configS3())

	key := "image/test/test_admin_medium.png"
	imageURL := "https://www.adobe.com/express/create/media_1fd4a6f2e688c9c269ec460ca1095e89953e33c28.png?width=2000&format=webply&optimize=medium"
	err := UploadImageToS3(uploader, key, imageURL)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println("errors ==>>", errorList)
}
