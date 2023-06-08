package model

import (
	"fmt"
	"myProject/datastore/postgres"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

const (
	queryInsert   = "INSERT INTO userdata(username, email, password) VALUES($1, $2, $3);"
	queryData = "SELECT * from userdata WHERE email = $1 AND password = $2"
)

func (s *User) Create() error {
	_, err := postgres.Db.Exec(queryInsert, s.Username, s.Email, s.Password)
	return err
}
func (s *User) CheckIn() error {
	//_ means ignore the variable
	_,err := postgres.Db.Query(queryData, s.Email, s.Password)
	fmt.Println(err)
	//row = nil
	if err != nil {
		//err = error
		return err
	}
	//if it is successful return nil
	return nil
}
