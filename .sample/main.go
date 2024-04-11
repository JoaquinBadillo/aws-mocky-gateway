package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type JSONResponse struct {
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Message string            `json:"message"`
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (*JSONResponse, error) {
	return &JSONResponse{
		Status:  200,
		Headers: map[string]string{"Content-Type": "application/json"},
		Message: "Hello Gateways!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
