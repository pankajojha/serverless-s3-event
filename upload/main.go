package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//GUID - generates a unique identifier
func GUID() (guid string) {
	tm := time.Now().UTC()
	t := tm.UnixNano() / 1000000
	fileDate := strconv.Itoa(tm.Year()) + "." + tm.Month().String() + "." + strconv.Itoa(tm.Day()) + "." + strconv.Itoa(tm.Hour()) + "." + strconv.Itoa(tm.Minute()) + "." + strconv.Itoa(tm.Second()) + "." + strconv.FormatInt(t, 10)
	guid = fileDate + ".json"
	return
}

//Request header key ...
const XAutherizationKey = "X-Notification-Secret"

// Request header value...
var XAutherizationValue = os.Getenv("XAutherization")

// Region needs be set ...
var REGION = os.Getenv("REGION")

// bucker name
var bucketName = os.Getenv("BUCKET")

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//////////////////
	// REGION = "us-east-1"
	// bucketName = "pci-1"
	// XAutherizationValue = "abcd1234"
	//////////////////

	fmt.Printf("Body size = %d.\n", len(request.Body))
	fmt.Println(" XAutherizationValue ", XAutherizationValue, " Region: ", REGION, " bucketName: ", bucketName)

	inputJSON := request.Body
	fmt.Println(" inputJSON ", inputJSON)

	isAuthenticated := false

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)

		if strings.EqualFold(key, XAutherizationKey) && strings.EqualFold(value, XAutherizationValue) {
			isAuthenticated = true
		}
	}

	if !isAuthenticated {
		fmt.Println(" Not Authenticated ... ")
		return events.APIGatewayProxyResponse{Body: "{status:420, success:false, reason : 'You are not autherized'}", StatusCode: 420}, nil
	}

	reader := strings.NewReader(inputJSON)
	newFileName := GUID()

	bucket := flag.String("bucket", bucketName, "The s3 bucket to upload to")

	fmt.Println("Received body: ", inputJSON, bucket, "fileName....", newFileName, " bucket: ", bucketName)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(REGION)},
	)

	if err != nil {
		fmt.Println("session-erro", err)
	}

	uploader := s3manager.NewUploader(sess)

	fmt.Println("uploader: ", uploader)

	//key := filepath.Base(file.Name())

	result, err1 := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucket),
		Key:    aws.String(newFileName),
		Body:   reader,
	})

	if err1 != nil {
		fmt.Println(err)
	}

	fmt.Printf("file uploaded to", result)

	return events.APIGatewayProxyResponse{Body: "{status:200, success:true}", StatusCode: 200}, nil

}

func main() {
	lambda.Start(handler)
}
