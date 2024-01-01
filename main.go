package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/LucioSchiavoni/go-aws/awsgo"
	"github.com/LucioSchiavoni/go-aws/db"
	"github.com/LucioSchiavoni/go-aws/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"os"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) events.CognitoEventUserPoolsPostConfirmation {

	awsgo.InicializoAws()

	if !ValidoParametros() {
		fmt.Println("Error en las credenciasles, debe enviar el secrect name")
		err := errors.New("Error en las credenciales, debe enviar el secret name")
		return event, err
	}
	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}
	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el secret " + err.Error())
		return event, err
	}
	err = db.SignUp(datos)
	return event, err

}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
