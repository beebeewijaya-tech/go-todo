package sql

import "context"

type Queries interface {
	// Users
	CreateUser(ctx context.Context, args CreateUserArgs) (User, error)
	GetUser(ctx context.Context, email string) (User, error)

	// Todos
	CreateTodo(ctx context.Context, args CreateTodoArgs) (Todo, error)
	GetTodo(ctx context.Context, id int64) (Todo, error)
	GetTodos(ctx context.Context, args GetTodosArgs) ([]Todo, error)
	UpdateTodo(ctx context.Context, args UpdateTodoArgs) (Todo, error)
	DeleteTodo(ctx context.Context, args DeleteTodoArgs) error
}

var _ Queries = (*Store)(nil)
