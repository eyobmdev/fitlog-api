package api

import (
	"log"
	"errors"
	"regexp"
	"github.com/eyobmdev/femProject/internal/store"
)

type registerUserRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Bio string `json:"bio"`
	Password string `json:"password"`
}

type UserHandler struct {
	userStore store.UserStore
	logger *log.Logger
}


func NewUserHandler(userStore store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		logger: logger,
	}
}

func (h *UserHandler) validateRegisterRequest(req *registerUserRequest) error {
	if req.Username == "" {
		return errors.New("Username is required")
	}
	if len(req.Username) > 50 {
		return errors.New("username cannot be greater than 50 characters")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email){
		return errors.New("ivalid email format")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}
	return nil
	
}
