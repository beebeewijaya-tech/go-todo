package api

import (
	"net/http"
	"time"

	"beebeewijaya.com/db/sql"
	"beebeewijaya.com/token"
	"beebeewijaya.com/util"
	"github.com/gin-gonic/gin"
)

type createUserParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Fullname string `json:"fullname" binding:"required"`
}

type createUserResponse struct {
	ID        int64
	Email     string
	Fullname  string
	CreatedAt time.Time
}

type loginUserParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginUserResponse struct {
	Email string
	Token string
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	hashed, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	reqUser := sql.CreateUserArgs{
		Email:    req.Email,
		Password: hashed,
		Fullname: req.Fullname,
	}

	user, err := s.db.CreateUser(ctx, reqUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	resp := createUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Fullname:  user.Fullname,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (s *Server) loginUser(ctx *gin.Context) {
	var req loginUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	user, err := s.db.GetUser(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	// Remove handling already handled at server.go
	maker, _ := token.NewMaker(s.config.GetString("JWT.SECRETKEY"))

	t, err := maker.GenerateToken(user.ID, user.Email, time.Minute*15)
	if err != nil {
		ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
		return
	}

	resp := loginUserResponse{
		Email: user.Email,
		Token: t,
	}

	ctx.JSON(http.StatusOK, resp)
}
