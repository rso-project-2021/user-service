package models

import (
	"context"
	"user-service/db"
)

type User struct {
	ID        int64  `json:"user_id" db:"user_id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type CreateUserParam struct {
	Username string
	Password string
	Email    string
}

type UpdateUserParam struct {
	Username string
	Password string
	Email    string
}

type ListUserParam struct {
	Offset int32
	Limit  int32
}

func (u User) GetByID(ctx context.Context, id int64) (user User, err error) {
	db := db.GetDB()

	const query = `SELECT * FROM "users" WHERE "user_id" = $1`
	err = db.GetContext(ctx, &user, query, id)

	return
}

func (u User) GetAll(ctx context.Context, arg ListUserParam) (users []User, err error) {
	db := db.GetDB()

	const query = `SELECT * FROM "users" OFFSET $1 LIMIT $2`
	users = []User{}
	err = db.SelectContext(ctx, &users, query, arg.Offset, arg.Limit)

	return
}

func (u User) Create(ctx context.Context, arg CreateUserParam) (User, error) {
	db := db.GetDB()

	const query = `
	INSERT INTO "users"("username", "password", "email") 
	VALUES ($1, $2, $3)
	RETURNING "user_id", "username", "password", "email", "created_at"
	`
	row := db.QueryRowContext(ctx, query, arg.Username, arg.Password, arg.Email)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)

	return user, err
}

func (u User) Update(ctx context.Context, arg UpdateUserParam, id int64) (User, error) {
	db := db.GetDB()

	const query = `
	UPDATE "users"
	SET "username" = $2,
		"password" = $3,
		"email"    = $4
	WHERE "user_id" = $1
	RETURNING "user_id", "username", "password", "email", "created_at"
	`
	row := db.QueryRowContext(ctx, query, id, arg.Username, arg.Password, arg.Email)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)

	return user, err
}

func (u User) Delete(ctx context.Context, id int64) error {
	db := db.GetDB()

	const query = `
	DELETE FROM users
	WHERE "user_id" = $1
	`
	_, err := db.ExecContext(ctx, query, id)

	return err
}
