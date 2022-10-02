package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/luckyparakh/goBank/db/sqlc"
	"github.com/luckyparakh/goBank/token"
	"github.com/luckyparakh/goBank/utils"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	// tokenMaker, err := token.NewJWTMaker(config.SYM_KEY)
	tokenMaker, err := token.NewPasetoMaker(config.SYM_KEY)
	if err != nil {
		return nil, fmt.Errorf("cannot create token")
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}
func (s *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", s.createUser)
	router.POST("/users/login", s.loginUser)
	
	authRouter := router.Group("/").Use(authMiddleware(s.tokenMaker))
	authRouter.POST("/accounts", s.createAccount)
	authRouter.GET("/accounts/:id", s.getAccount)
	authRouter.GET("/accounts/", s.listAccounts)
	authRouter.POST("/transfer", s.createTransfer)

	s.router = router
}

func (s *Server) StartServer(addr string) error {
	return s.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
