package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"rwa/internal/middleware"
	request "rwa/internal/transport/request/user"
	"rwa/internal/transport/resource"
	usecase "rwa/internal/usecase/user"
	"rwa/pkg/response"
	"rwa/pkg/validator"
	"strings"
)

type UsersController struct {
	UseCase usecase.UseCase
}

func NewUserController(useCase usecase.UseCase) *UsersController {
	return &UsersController{
		UseCase: useCase,
	}
}

func (c *UsersController) Register(router *mux.Router) {
	router.HandleFunc("/api/user", middleware.ErrorHandler(c.Create)).Methods("POST")
	router.HandleFunc("/api/user/login", middleware.ErrorHandler(c.Login)).Methods("POST")
	router.HandleFunc("/api/user", middleware.ErrorHandler(c.Authenticate)).Methods("GET")
}

func (c *UsersController) Authenticate(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Token ")
	fmt.Println(token)
	loggedUser, err := c.UseCase.Authenticate(ctx, token)
	if err != nil {
		return err
	}

	return response.JSON(w, http.StatusOK, resource.ConvertDomainToResource(*loggedUser))
}

func (c *UsersController) Login(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	body := r.Body

	defer r.Body.Close()

	bodyByte, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("byte %s", err)
	}
	var usersWrapper request.UsersWrapper
	err = json.Unmarshal(bodyByte, &usersWrapper)
	if err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	err = validator.Struct(usersWrapper.Users)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	domainUser := request.ConvertUserToDomain(usersWrapper.Users)

	loggedUser, err := c.UseCase.Login(ctx, &domainUser)

	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}
	w.Header().Set("Authorization", "Token "+loggedUser.Token)

	return response.JSON(w, http.StatusOK, resource.ConvertDomainToResource(*loggedUser))
}

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) error {

	ctx := r.Context()
	body := r.Body

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(body)

	bodyByte, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("byte %s", err)
	}
	var usersWrapper request.UsersWrapper
	err = json.Unmarshal(bodyByte, &usersWrapper)
	if err != nil {
		return err
	}

	err = validator.Struct(usersWrapper.Users)
	if err != nil {
		return err
	}

	user := request.ConvertUserToDomain(usersWrapper.Users)

	err = c.UseCase.Create(ctx, &user)

	if err != nil {
		return err
	}

	return response.JSON(w, http.StatusCreated, resource.ConvertDomainToResource(user))
}
