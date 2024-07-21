package models

import (
	"errors"

	"eventbooking.com/rest-api/db"
	"eventbooking.com/rest-api/utils"
)

type User struct {
	ID        int64
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	Full_name string
}

func (this User) Save() error {
	query := "INSERT INTO users(email, password, full_name) VALUES (?, ?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedpassword, err := utils.HashPassword(this.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(this.Email, hashedpassword, this.Full_name)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	this.ID = userId

	return err
}

func (this *User) ValidateCreds() error {
	query := "SELECT id, password FROM users where email = ?"
	row := db.DB.QueryRow(query, this.Email)

	var retrievedPassword string

	err := row.Scan(&this.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Invalid Credentials")
	}

	isPassword := utils.CheckPassword(this.Password, retrievedPassword)

	if !isPassword {
		return errors.New("Invalid Credentials")
	}

	return nil
}
