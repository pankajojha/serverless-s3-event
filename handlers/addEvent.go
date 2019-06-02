// src/handlers/addEvents.go

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"./util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var ddb *dynamodb.DynamoDB
var uploader *s3manager.Uploader

//GUID - generates a unique identifier
func GUID() (guid string) {
	tm := time.Now().UTC()
	t := tm.UnixNano() / 1000000
	fileDate := strconv.Itoa(tm.Year()) + "." + tm.Month().String() + "." + strconv.Itoa(tm.Day()) + "." + strconv.Itoa(tm.Hour()) + "." + strconv.Itoa(tm.Minute()) + "." + strconv.Itoa(tm.Second()) + "." + strconv.FormatInt(t, 10)
	guid = fileDate
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

func init() {
	REGION := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &REGION,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) // Create DynamoDB client
		uploader = s3manager.NewUploader(session)

	}
}

func AddEvent(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

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
	fmt.Println(" Authenticated ... ", isAuthenticated)
	if !isAuthenticated {
		return events.APIGatewayProxyResponse{Body: "{status:420, success:false, reason : 'You are not autherized'}", StatusCode: 420}, nil
	}

	reader := strings.NewReader(inputJSON)
	idFileName := GUID()
	newFileName := idFileName + ".json"

	bucket := flag.String("bucket", bucketName, "The s3 bucket to upload to")
	fmt.Println("Received body: ", inputJSON, bucket, "fileName....", newFileName, " bucket: ", bucketName)

	//uploader := s3manager.Uploader(session)

	result, err1 := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucket),
		Key:    aws.String(newFileName),
		Body:   reader,
	})

	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("file uploaded to", result)

	var (
		id        = idFileName
		tableName = aws.String(os.Getenv("MC_TABLE_NAME"))
	)

	// Initialize todo
	todo := &util.Event{
		ID:        id,
		Done:      false,
		CreatedAt: time.Now().String(),
	}

	// Parse request body
	json.Unmarshal([]byte(request.Body), todo)

	// Write to DynamoDB
	item, _ := dynamodbattribute.MarshalMap(todo)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tableName,
	}
	if _, err := ddb.PutItem(input); err != nil {
		return events.APIGatewayProxyResponse{ // Error HTTP response
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	} else {
		body, _ := json.Marshal(todo)
		return events.APIGatewayProxyResponse{ // Success HTTP response
			Body:       string(body),
			StatusCode: 200,
		}, nil
	}

}

func main() {
	lambda.Start(AddEvent)
}
