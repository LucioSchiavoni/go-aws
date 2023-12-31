package secret

import (
	// "encoding/json"
	"fmt"
	// "github.com/LucioSchiavoni/go-aws/aws"
	"github.com/LucioSchiavoni/go-aws/models"
	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println(" =>" + nombreSecret)
	svc := secretsmanager.NewFromConfig(aws.Cfg)
	clave, err := svc.GetSecretValue()
}
