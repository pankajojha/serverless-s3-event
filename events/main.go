package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, s3Event events.S3Event) {

	fmt.Println(" Received s3 event ", s3Event)

	data, _ := json.Marshal(s3Event)
	//Now convert to a string and output
	//Cloudwatch picks up the json and formats it nicely for us. :)
	streventinfo := string(data)

	// stdout and stderr are sent to AWS CloudWatch Logs
	fmt.Printf("S3 Event : %s\n", streventinfo)

	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
	}
}

func main() {
	lambda.Start(handleRequest)
}
