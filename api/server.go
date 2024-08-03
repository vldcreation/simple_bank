package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	db "github.com/vldcreation/simple_bank/db/sql/postgresql/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)

	// user routes
	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

type apiError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func errorResponse(err error) gin.H {
	var ve validator.ValidationErrors
	if ok := errors.As(err, &ve); ok {
		_errs := make([]apiError, 0, len(ve))
		for _, fieldErr := range ve {
			_errs = append(_errs, msgFieldErrResponse(fieldErr))
		}
		return gin.H{"errors": _errs}
	}

	if pqErr, ok := err.(*pq.Error); ok {
		return msgDBErrResponse(*pqErr)
	}

	return gin.H{"error": err.Error()}
}

func msgDBErrResponse(err pq.Error) gin.H {
	switch err.Code {
	case "23505", "unique_violation":
		return gin.H{"error": "already exists"}
	case "23503", "foreign_key_violation":
		return gin.H{"error": "does not exist"}
	default:
		return gin.H{"error": "internal server error"}
	}
}

func msgFieldErrResponse(field validator.FieldError) apiError {
	err := apiError{
		Field: field.Field(),
	}
	tag := field.Tag()

	switch tag {
	case "required":
		err.Msg = "should be not empty"
	case "email":
		err.Msg = "should be a valid email address"
	case "min":
		err.Msg = "must be at least " + field.Param() + " characters long"
	case "max":
		err.Msg = "must be at most " + field.Param() + " characters long"
	case "alphanum":
		err.Msg = "must be alphanumeric"
	default:
		err.Msg = "invalid input"
	}

	return err
}
