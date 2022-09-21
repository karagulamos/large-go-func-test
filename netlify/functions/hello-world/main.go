package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	content := fmt.Sprintf("%s", "TEXT_TO_REPLACE")
	
	// should appear in console
	fmt.Printf("Hello, %s!", humanize(len(content)))

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            fmt.Sprintf("Hello, %s!", humanize(len(content))),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func humanize(size int) string {
	switch {
	case size < 1024 :
		return fmt.Sprintf("%d bytes", size)
	case size < 1024 * 1024:
		return fmt.Sprintf("%.2f kilobytes", float64(size) / 1024)
	default:
		return fmt.Sprintf("%.2f megabytes", float64(size) / (1024 * 1024))
	}
}
