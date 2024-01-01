package secret

import (
	"encoding/json"
	"fmt"

	"github.com/LucioSchiavoni/go-aws/awsgo"
	"github.com/LucioSchiavoni/go-aws/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println(" =>" + nombreSecret)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" =>Lectura secret OK: " + nombreSecret)
	return datosSecret, nil
}
