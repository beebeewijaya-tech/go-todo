package util

import "errors"

const (
	Authorization  = "Authorization"
	Bearer         = "Bearer"
	AuthPayloadKey = "AuthorizationPayload"
)

var ErrEmptyToken = errors.New("token is empty")
var ErrTokenInvalid = errors.New("token is invalid")
var ErrTokenExpired = errors.New("token expires")
var ErrAuthForbidden = errors.New("auth forbidden to execute logic")
var ErrEmptyRow = errors.New("row is empty")

var SuccessDeleteTodo = "Successfully deleting todo"
