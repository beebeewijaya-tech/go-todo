package api

import (
	"net/http"

	"beebeewijaya.com/db/sql"
	"beebeewijaya.com/util"
	"github.com/gin-gonic/gin"
)

type createTodoBody struct {
	Title       string           `json:"title" binding:"required"`
	Description string           `json:"description" binding:"required"`
	Priority    sql.PriorityType `json:"priority" binding:"required,min=1,max=3"`
}

type getTodosQuery struct {
	PageSize int64 `form:"page_size" binding:"min=5"`
	Page     int64 `form:"page" binding:"min=1"`
}

type getTodoParams struct {
	ID int64 `uri:"id" binding:"required,min=0"`
}

type updateTodoArgs struct {
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	Priority    sql.PriorityType `json:"priority,omitempty" binding:"min=1,max=3"`
}

type deleteTodoResp struct {
	ID      int64
	Message string
}

func (s *Server) createTodo(ctx *gin.Context) {
	var req createTodoBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload, err := GetAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	t := sql.CreateTodoArgs{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Author:      authPayload.Username,
	}

	todo, err := s.db.CreateTodo(ctx, t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (s *Server) getTodos(ctx *gin.Context) {
	var req getTodosQuery
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload, err := GetAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	t := sql.GetTodosArgs{
		Page:     (req.Page - 1) * req.PageSize,
		PageSize: req.PageSize,
		Author:   authPayload.Username,
	}

	todos, err := s.db.GetTodos(ctx, t)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(util.ErrTokenInvalid))
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (s *Server) getTodo(ctx *gin.Context) {
	var req getTodoParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	todo, err := s.db.GetTodo(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (s *Server) updateTodo(ctx *gin.Context) {
	var params getTodoParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	var req updateTodoArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload, err := GetAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	t, err := s.db.GetTodo(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	if authPayload.Username != t.Author {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(util.ErrAuthForbidden))
		return
	}

	title := req.Title
	if title == "" {
		title = t.Title
	}

	desc := req.Description
	if desc == "" {
		desc = t.Description
	}

	priority := req.Priority
	if priority == 0 {
		priority = t.Priority
	}

	ut := sql.UpdateTodoArgs{
		ID:          t.ID,
		Title:       title,
		Description: desc,
		Priority:    priority,
	}
	todo, err := s.db.UpdateTodo(ctx, ut)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (s *Server) deleteTodo(ctx *gin.Context) {
	var req getTodoParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload, err := GetAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	dt := sql.DeleteTodoArgs{
		ID:     req.ID,
		Author: authPayload.Username,
	}

	err = s.db.DeleteTodo(ctx, dt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	resp := deleteTodoResp{
		ID:      req.ID,
		Message: util.SuccessDeleteTodo,
	}

	ctx.JSON(http.StatusOK, resp)
}
