package db

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
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

func (store *Store) GetUserByID(ctx context.Context, id int64) (user User, err error) {

	const query = `SELECT * FROM "users" WHERE "user_id" = $1`
	err = store.db.GetContext(ctx, &user, query, id)

	return
}

func (store *Store) GetAllUsers(ctx context.Context, arg ListUserParam) (users []User, err error) {

	const query = `SELECT * FROM "users" OFFSET $1 LIMIT $2`
	users = []User{}
	err = store.db.SelectContext(ctx, &users, query, arg.Offset, arg.Limit)

	return
}

func (store *Store) CreateUser(ctx context.Context, arg CreateUserParam) (User, error) {

	const query = `
	INSERT INTO "users"("username", "password", "email") 
	VALUES ($1, $2, $3)
	RETURNING "user_id", "username", "password", "email", "created_at"
	`
	row := store.db.QueryRowContext(ctx, query, arg.Username, arg.Password, arg.Email)

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

func (store *Store) UpdateUser(ctx context.Context, arg UpdateUserParam, id int64) (User, error) {

	const query = `
	UPDATE "users"
	SET "username" = $2,
		"password" = $3,
		"email"    = $4
	WHERE "user_id" = $1
	RETURNING "user_id", "username", "password", "email", "created_at"
	`
	row := store.db.QueryRowContext(ctx, query, id, arg.Username, arg.Password, arg.Email)

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

func (store *Store) DeleteUser(ctx context.Context, id int64) error {

	const query = `
	DELETE FROM users
	WHERE "user_id" = $1
	`
	_, err := store.db.ExecContext(ctx, query, id)

	return err
}

func (store *Store) GetUserByUsername(ctx context.Context, arg CreateUserParam) (User, error) {

	const query = `SELECT * FROM "users" WHERE "username" = $1`
	row := store.db.QueryRowContext(ctx, query, arg.Username)

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
