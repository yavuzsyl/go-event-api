package models

import "eventapi/db"

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

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.SetId(userId)
	return err
}
