package models

import (
	"Uber/config"
	"fmt"
)

type Signup struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

func CreateNewUser(newuser Signup) (*Signup, error) {
	var err error
	if newuser.Name == "" || newuser.Username == "" || newuser.Password == "" || newuser.Mobile == "" {
		fmt.Println("Error")
	}
	db := config.GetDB()
	sqlStatement := `INSERT INTO users (name, username, password, mobile)
						VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, newuser.Name, newuser.Username, newuser.Password, newuser.Mobile)

	if err != nil {
		panic(err)
	}
	return &newuser, nil
}
