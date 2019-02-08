package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Post struct {
	Field string `json:"field"`
	Pk    string `json:"pk"`
	Sk    string `json:"sk"`
	Title string `json:"title"`
}

type Response struct {
	Pk    string `json:"pk"`
	Sk    string `json:"sk"`
	Title string `json:"title"`
}

func hander(arg Post) (Response, error) {

	response := Response{}

	fmt.Println("[DEBUG]field : " + arg.Field)
	fmt.Println("[DEBUG]pk : " + arg.Pk)
	fmt.Println("[DEBUG]sk : " + arg.Sk)

	fmt.Println("[DEBUG]Start create session.")
	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		return Response{}, err
	}
	svc := dynamodb.New(session)

	switch arg.Field {
	case "putPost":
		fmt.Println("[DEBUG]start putPost process")
		input := &dynamodb.PutItemInput{
			Item: map[string]*dynamodb.AttributeValue{
				"pk": {
					S: aws.String(arg.Pk),
				},
				"sk": {
					S: aws.String(arg.Sk),
				},
				"title": {
					S: aws.String(arg.Title),
				},
			},
			TableName: aws.String("appsync-lambda-go"),
		}

		fmt.Println("[DEBUG]Start put item to DynamoDB.")

		_, err = svc.PutItem(input)
		if err != nil {
			return Response{}, err
		}

		response.Pk = arg.Pk
		response.Sk = arg.Sk
		response.Title = arg.Title

	case "singlePost":
		fmt.Println("[DEBUG]start singlePost process")
		input := &dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"pk": {
					S: aws.String(arg.Pk),
				},
				"sk": {
					S: aws.String(arg.Sk),
				},
			},
			TableName: aws.String("appsync-lambda-go"),
		}

		fmt.Println("[DEBUG]Start get item from DynamoDB.")

		result, err := svc.GetItem(input)
		if err != nil {
			return Response{}, err
		}

		resultData := &Response{}
		if err := dynamodbattribute.UnmarshalMap(result.Item, resultData); err != nil {
			return Response{}, err
		}
		response.Pk = resultData.Pk
		response.Sk = resultData.Sk
		response.Title = resultData.Title
	}

	return response, nil
}

func main() {
	lambda.Start(hander)
}
