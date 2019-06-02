package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var WebhookUrl = os.Getenv("webhook_post_url")

func handleRequest(ctx context.Context, s3Event events.S3Event) {

	fmt.Println(" Received s3 event ", s3Event, WebhookUrl)

	url := WebhookUrl
	fmt.Println("URL:>", url)

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

	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	var jsonStr = []byte(streventinfo)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "hal-header")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func main() {
	lambda.Start(handleRequest)
}
