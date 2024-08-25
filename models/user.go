package models

import (
	"errors"
	"fmt"

	db "example.com/gin-go-api/sql-db"
	"example.com/gin-go-api/utils"
)

type User struct {
	ID       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) SaveToQL() error {
	query := `INSERt INTO users(email, password) VALUES (?,?)`
	stmt, err := db.SQL_DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println("HassPassWORD ERR ->>", err)
		return err
	}
	result, err := stmt.Exec(u.Email, hashPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId

	return err

}

func (u *User) ValidateCreds() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.SQL_DB.QueryRow(query, u.Email)

	var retrivePassword string
	err := row.Scan(&u.ID, &retrivePassword)
	if err != nil {
		fmt.Println("VALIDATE CREDS ->>", err)
		return err
	}

	passwordIsValidate := utils.CheckPasswordHash(u.Password, retrivePassword)

	if !passwordIsValidate {
		return errors.New("Creds are invalid->>>")
	}

	return nil
}
