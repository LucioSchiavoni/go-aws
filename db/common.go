package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LucioSchiavoni/go-aws/models"
	"github.com/LucioSchiavoni/go-aws/secret"
	_ "github.com/go-sql-driver/mysq"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

// es de tipo puntero * por un tema de velocidad

func ReadSecret() error {
	SecretModel, err = secret.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion exitosa de la DB")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	return dsn
}
