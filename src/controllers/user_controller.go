package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-golang-jwt/src/entity"
	"rest-api-golang-jwt/src/helpers"
	"rest-api-golang-jwt/src/services"
	"strings"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c UserController) Login(x http.ResponseWriter, r *http.Request) {
	x.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var loginRequest entity.UserLoginRequest
		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := c.UserService.Login(r.Context(), loginRequest)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		responseJson, err := json.Marshal(user)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		x.Write(responseJson)
		return
	}
}

func (c UserController) Get(x http.ResponseWriter, r *http.Request) {
	x.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		users, err := c.UserService.Get(r.Context())
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		response := Response{
			Code: http.StatusOK,
			Data: users,
		}
		responseJson, err := json.Marshal(response)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		x.Write(responseJson)
		return
	}
	http.Error(x, "", http.StatusBadRequest)
}

func (c UserController) Insert(x http.ResponseWriter, r *http.Request) {
	x.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var user entity.UserRequest
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := c.UserService.Insert(r.Context(), user)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		response := Response{
			Code: http.StatusOK,
			Data: result,
		}
		responseJson, err := json.Marshal(response)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		x.Write(responseJson)
		return
	}
	http.Error(x, "", http.StatusBadRequest)
}

func (c UserController) GetById(x http.ResponseWriter, r *http.Request) {
	x.Header().Set("Content-Type", "application/json")

	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		http.Error(x, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	bool, err := helpers.ValidateToken(tokenString)
	if err != nil {
		http.Error(x, err.Error(), http.StatusInternalServerError)
		return
	}
	if !bool {
		http.Error(x, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method == "GET" {
		id, err := helpers.ExtractToken(tokenString)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := c.UserService.GetById(r.Context(), id)

		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		response := Response{
			Code: http.StatusOK,
			Data: user,
		}
		responseJson, err := json.Marshal(response)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
			return
		}
		x.Write(responseJson)
		return
	}
	http.Error(x, "", http.StatusBadRequest)
}

type Response struct {
	Code  int
	Data  interface{}
	Error string
}
