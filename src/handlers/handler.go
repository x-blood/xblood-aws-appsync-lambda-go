package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Post struct {
	Pk string `json:"pk"`
	Sk string `json:"sk"`
}

func hander(arg Post) (Post, error) {

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(arg.Pk),
			},
			"sk": {
				S: aws.String(arg.Sk),
			},
		},
		TableName: aws.String("post"),
	}

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		return Post{}, err
	}

	svc := dynamodb.New(session)

	_, err = svc.PutItem(input)

	return arg, nil
}

func main() {
	lambda.Start(hander)
}
