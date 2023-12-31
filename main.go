package main

import (
	"context"
	"github.com/LucioSchiavoni/go-aws/aws"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) events.CognitoEventUserPoolsPostConfirmation {
	aws.InicializoAws()
}
