package models

import (
	"eventapi/db"
	"eventapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (e *User) SetId(id int64) {
	e.ID = id
}

func (u User) Save() error {
	command := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(command)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.SetId(userId)
	return err
}
