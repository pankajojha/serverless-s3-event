package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"./util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type ListEventsResponse struct {
	Events []util.Event `json:"todos"`
}

//Request header key ...
const XAutherizationKey = "X-Notification-Secret"

// Request header value...
var XAutherizationValue = os.Getenv("XAutherization")

var ddb *dynamodb.DynamoDB

func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) // Create DynamoDB client
	}
}

func ListEvents(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("ListTodos")

	// if !util.IsAuthenticated(request.Headers) {
	// 	return events.APIGatewayProxyResponse{Body: "{status:420, success:false, reason : 'You are not autherized'}", StatusCode: 420}, nil
	// }

	var (
		id        = request.PathParameters["id"]
		tableName = aws.String(os.Getenv("MC_TABLE_NAME"))
	)

	fmt.Println(" id TODO query ", id)

	// Read from DynamoDB
	input := &dynamodb.ScanInput{
		TableName: tableName,
	}
	result, _ := ddb.Scan(input)

	// Construct todos from response
	var todos []util.Event
	for _, i := range result.Items {
		todo := util.Event{}
		if err := dynamodbattribute.UnmarshalMap(i, &todo); err != nil {
			fmt.Println("Failed to unmarshal")
			fmt.Println(err)
		}
		todos = append(todos, todo)
	}

	// Success HTTP response
	body, _ := json.Marshal(&ListEventsResponse{
		Events: todos,
	})
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(ListEvents)
}
