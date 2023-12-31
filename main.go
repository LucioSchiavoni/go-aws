package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/LucioSchiavoni/go-aws/aws"
	"github.com/LucioSchiavoni/go-aws/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"os"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) events.CognitoEventUserPoolsPostConfirmation {

	aws.InicializoAws()

	if !ValidoParametros() {
		fmt.Println("Error en las credenciasles, debe enviar el secrect name")
		err := errors.New("Error en las credenciales, debe enviar el secret name")
		return event, err
	}
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
