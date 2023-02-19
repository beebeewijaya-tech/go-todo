package sql

import (
	"context"
	"time"

	"beebeewijaya.com/util"
)

const createTodo = `
INSERT INTO todos (
	title,
	description,
	author,
	priority
) VALUES (
	$1,
	$2,
	$3,
	$4
) RETURNING *
`

type CreateTodoArgs struct {
	Title       string
	Description string
	Author      int64
	Priority    PriorityType
}

func (s *Store) CreateTodo(ctx context.Context, args CreateTodoArgs) (Todo, error) {
	row := s.db.QueryRowContext(ctx, createTodo, args.Title, args.Description, args.Author, args.Priority)

	var t Todo
	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Priority,
		&t.Author,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	return t, err
}

const getTodo = `
SELECT * FROM todos
WHERE id = $1
`

func (s *Store) GetTodo(ctx context.Context, id int64) (Todo, error) {
	row := s.db.QueryRowContext(ctx, getTodo, id)

	var t Todo
	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Priority,
		&t.Author,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	return t, err
}

const getTodos = `
SELECT * FROM todos
WHERE author = $1
LIMIT $2
OFFSET $3
`

type GetTodosArgs struct {
	Author   int64
	Page     int64
	PageSize int64
}

func (s *Store) GetTodos(ctx context.Context, args GetTodosArgs) ([]Todo, error) {
	rows, err := s.db.QueryContext(ctx, getTodos, args.Author, args.PageSize, args.Page)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var t Todo
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Priority,
			&t.Author,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}

		todos = append(todos, t)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

const updateTodo = `
UPDATE todos
SET title = $2, description = $3, priority = $4, updated_at = $5
WHERE id = $1
RETURNING *
`

type UpdateTodoArgs struct {
	ID          int64
	Title       string
	Description string
	Priority    PriorityType
	UpdatedAt   time.Time
}

func (s *Store) UpdateTodo(ctx context.Context, args UpdateTodoArgs) (Todo, error) {
	row := s.db.QueryRowContext(ctx, updateTodo, args.ID, args.Title, args.Description, args.Priority, args.UpdatedAt)

	var t Todo
	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Priority,
		&t.Author,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	return t, err
}

const deleteTodo = `
DELETE FROM todos
WHERE id = $1 AND author = $2
`

type DeleteTodoArgs struct {
	ID     int64
	Author int64
}

func (s *Store) DeleteTodo(ctx context.Context, args DeleteTodoArgs) error {
	row, err := s.db.ExecContext(ctx, deleteTodo, args.ID, args.Author)
	if err != nil {
		return err
	}

	affected, err := row.RowsAffected()
	if affected < 1 || err != nil {
		return util.ErrEmptyRow
	}

	return nil
}
