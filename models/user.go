package models

import (
	"user-service/db"
)

type User struct {
	ID        int64  `json:"user_id" db:"user_id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

func (u User) GetByID(id string) (user User, err error) {
	db := db.GetDB()

	// get user by ID
	query := `SELECT * FROM "users" WHERE "user_id" = $1`
	err = db.Get(&user, query, id)

	if err != nil {
		panic(err)
	}

	return
}

func (u User) GetAll() (user []User, err error) {
	db := db.GetDB()

	// get list of users
	query := `SELECT * FROM "users"`
	err = db.Select(&user, query)

	if err != nil {
		panic(err)
	}

	return
}
