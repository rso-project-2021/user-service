package models

import (
	"user-service/db"
)

type Account struct {
	ID        int64  `json:"account_id" db:"account_id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

func (u Account) GetByID(id string) (*Account, error) {

	// get database connection
	db := db.GetDB()

	// get account by ID
	query := `SELECT * FROM "accounts" WHERE "account_id" = $1`
	account := Account{}
	err := db.Get(&account, query, id)

	// check if query threw an error
	if err != nil {
		panic(err)
	}

	return &account, nil
}
