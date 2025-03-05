package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)


type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error{
		query:= `
	INSERT INTO  users(email, password)
	VALUES(?,?)
	`
	stm, err := db.DB.Prepare(query)
	if err != nil {
    	return err
    }

	hashedPassword, err := utils.HashPassword(u.Password)
		if err != nil {
       return err
    }
	result, err := stm.Exec(u.Email,hashedPassword)
	defer stm.Close()

	if err != nil {
       return err
    }
	userId,err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query,u.Email)

	var retreivedPassword  string 

	err := row.Scan(&u.ID,&retreivedPassword)
	if err != nil {
        return err
    }
	passWordIsValid := utils.CheckPasswordHash(retreivedPassword,u.Password)

	if !passWordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
