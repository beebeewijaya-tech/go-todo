package sql

import (
	"context"
)

const createUser = `
INSERT INTO users (
	email,
	password,
	fullname
) VALUES (
	$1,
	$2,
	$3
) RETURNING *
`

type CreateUserArgs struct {
	Email    string
	Password string
	Fullname string
}

func (s *Store) CreateUser(ctx context.Context, args CreateUserArgs) (User, error) {
	row := s.db.QueryRowContext(ctx, createUser, args.Email, args.Password, args.Fullname)

	var u User
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Fullname,
		&u.Password,
		&u.CreatedAt,
	)

	return u, err
}

const getUser = `
SELECT * FROM users
WHERE email = $1
`

func (s *Store) GetUser(ctx context.Context, email string) (User, error) {
	row := s.db.QueryRowContext(ctx, getUser, email)

	var u User
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Fullname,
		&u.Password,
		&u.CreatedAt,
	)

	return u, err
}
