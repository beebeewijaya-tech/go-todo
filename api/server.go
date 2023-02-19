package api

import (
	"log"

	"beebeewijaya.com/db/sql"
	"beebeewijaya.com/middleware"
	"beebeewijaya.com/token"
	"beebeewijaya.com/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server struct {
	router     *gin.Engine
	config     *viper.Viper
	db         sql.Queries
	tokenMaker token.Maker
}

func NewServer(db sql.Queries, c *viper.Viper) *Server {
	s := &Server{}

	maker, err := token.NewMaker(c.GetString("JWT.SECRETKEY"))
	if err != nil {
		log.Fatalf("error when initting JWT token: %v", err)
	}

	s.db = db
	s.config = c
	s.tokenMaker = maker
	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	router := gin.Default()

	// Users router
	router.POST("/user/create", s.createUser)
	router.POST("/user/login", s.loginUser)

	authRoutes := router.Group("/").Use(middleware.JWTAuthMiddleware(s.tokenMaker))

	// Todos router
	authRoutes.GET("/todo/:id", s.getTodo)
	authRoutes.GET("/todos", s.getTodos)
	authRoutes.POST("/todo", s.createTodo)
	authRoutes.PUT("/todo/:id", s.updateTodo)
	authRoutes.DELETE("/todo/:id", s.deleteTodo)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func GetAuthPayload(ctx *gin.Context) (*token.Payload, error) {
	authPayload, ok := ctx.MustGet(util.AuthPayloadKey).(*token.Payload)
	if !ok {
		return nil, util.ErrTokenInvalid
	}

	return authPayload, nil
}
