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
const XAutherizationKey = "X-Autherization"

// Request header value...
var XAutherizationValue = os.Getenv("XAutherization")

// Region needs be set ...
var REGION = os.Getenv("AWS_REGION")

var buketName = os.Getenv("bucketName")

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("Body size = %d.\n", len(request.Body))

	isAuthenticated := false

	if XAutherizationValue == "" {
		fmt.Println(" set the XAutherization value")
		//XAutherizationValue = "Test123"
	}

	if REGION == "" {
		fmt.Println(" Region value is not set please, trying us-east-1 ")
		REGION = "us-east-1"
	}

	if buketName == "" {
		fmt.Println(" bucketName needs to be defined")
	}

	fmt.Println(" XAutherizationValue ", XAutherizationValue, " Region: ", REGION, " bucketName: ", bucketName)

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)

		if strings.EqualFold(key, XAutherizationKey) && strings.EqualFold(value, XAutherizationValue) {
			isAuthenticated = true
		}
	}

	if !isAuthenticated {
		return events.APIGatewayProxyResponse{Body: "{status:420, success:false, reason : 'You are not autherized'}", StatusCode: 420}, nil
	}

	inputJSON := request.Body
	reader := strings.NewReader(inputJSON)

	newFileName := GUID()

	bucket := flag.String("bucket", bucketName, "The s3 bucket to upload to")
	//filename := flag.String(newFileName, "", "The file to be uploaded to s3")

	fmt.Println("Received body: ", inputJSON, bucket, "fileName....", newFileName)

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
