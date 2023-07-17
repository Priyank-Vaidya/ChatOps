package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Priyank-Vaidya/ChatOps/routers"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading ENV")
	}

	slackToken := os.Getenv("SLACK_AUTH_TOKEN")
	log.Println(slackToken)

	router := gin.Default()
	routers.InitializeRoutes(router)

	// Use gin adapter to convert the request and response types
	resp, err := ginadapter.New(router).Proxy(request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	// Return the response from the Gin router
	return events.APIGatewayProxyResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       resp.Body.String(),
	}, nil
}

func main() {
	// Start the Lambda handler
	lambda.Start(handler)
}