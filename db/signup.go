package db

import (
	"fmt"

	"github.com/LucioSchiavoni/go-aws/models"
	"github.com/LucioSchiavoni/go-aws/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza registro")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()
	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "', '" + sig.UserUUID + "', '" + tools.FechaMySQL() + "')"
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
