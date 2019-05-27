package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {

	tm := time.Now().UTC()
	t := tm.UnixNano() / 1000000
	fileDate := strconv.Itoa(tm.Year()) + "." + tm.Month().String() + "." + strconv.Itoa(tm.Day()) + "." + strconv.Itoa(tm.Hour()) + "." + strconv.Itoa(tm.Minute()) + "." + strconv.Itoa(tm.Second()) + "." + strconv.FormatInt(t, 10)

	fmt.Println(fileDate)

	lambda.Start(Handler)
}
