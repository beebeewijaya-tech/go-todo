package sql

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"beebeewijaya.com/util"
	"github.com/stretchr/testify/require"
)

func createRandomTodo(t *testing.T, author int64) Todo {
	args := CreateTodoArgs{
		Title:       util.RandomString(12),
		Description: util.RandomString(100),
		Author:      util.RandomInt(1, 500),
		Priority:    PriorityType(util.RandomInt(0, 2)),
	}

	if author != 0 {
		args.Author = author
	}

	todo, err := testQuery.CreateTodo(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, todo.Title, args.Title)
	require.Equal(t, todo.Description, args.Description)
	require.Equal(t, todo.Author, args.Author)
	require.Equal(t, todo.Priority, args.Priority)
	require.NotZero(t, todo.ID)
	require.NotZero(t, todo.CreatedAt)
	require.NotZero(t, todo.UpdatedAt)

	return todo
}

func TestCreateTodo(t *testing.T) {
	createRandomTodo(t, 0)
}

func TestGetTodos(t *testing.T) {
	pageSize := 3
	page := 0
	author := 1

	args := GetTodosArgs{
		Page:     int64(page),
		PageSize: int64(pageSize),
		Author:   int64(author),
	}

	for i := 0; i < pageSize; i++ {
		createRandomTodo(t, int64(author))
	}

	todos, err := testQuery.GetTodos(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, todos, pageSize)

	for _, todo := range todos {
		require.NotZero(t, todo.ID)
		require.NotZero(t, todo.Title)
		require.NotZero(t, todo.Description)
		require.NotZero(t, todo.Author)
		require.NotZero(t, todo.CreatedAt)
		require.NotZero(t, todo.UpdatedAt)
	}
}

func TestUpdateTodo(t *testing.T) {
	todo := createRandomTodo(t, 0)

	args := UpdateTodoArgs{
		ID:          todo.ID,
		Title:       util.RandomString(12),
		Description: util.RandomString(100),
		Priority:    PriorityType(util.RandomInt(0, 2)),
		UpdatedAt:   time.Now(),
	}

	todo2, err := testQuery.UpdateTodo(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo2.ID, todo.ID)
	require.Equal(t, todo2.Title, args.Title)
	require.Equal(t, todo2.Description, args.Description)
	require.Equal(t, todo2.Priority, args.Priority)
	require.WithinDuration(t, todo2.CreatedAt, todo.CreatedAt, time.Second)
	require.WithinDuration(t, todo2.UpdatedAt, args.UpdatedAt, time.Second)
}

func TestDelete(t *testing.T) {
	todo := createRandomTodo(t, 0)

	args := DeleteTodoArgs{
		ID:     todo.ID,
		Author: todo.Author,
	}

	err := testQuery.DeleteTodo(context.Background(), args)
	require.NoError(t, err)

	todo2, err := testQuery.GetTodo(context.Background(), args.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, todo2)
}
