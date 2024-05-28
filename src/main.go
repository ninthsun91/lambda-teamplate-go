package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct{}

func HandleRequest(ctx context.Context, event *LambdaEvent) (*string, error) {
	if event != nil {
		return nil, fmt.Errorf("unexpected event data")
	}

	msg := "Hello, World!"
	return &msg, nil
}

func main() {
	lambda.Start(HandleRequest)
}
