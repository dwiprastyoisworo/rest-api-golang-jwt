package main

import (
	"fmt"
	"net/http"
	"rest-api-golang-jwt/src/config"
	"rest-api-golang-jwt/src/controllers"
	"rest-api-golang-jwt/src/repositories"
	"rest-api-golang-jwt/src/services"
)

func main() {
	db, err := config.PostgresConnect()
	if err != nil {
		panic(err)
	}
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db)
	userController := controllers.NewUserController(userService)

	http.HandleFunc("/users", userController.Get)
	http.HandleFunc("/users/create", userController.Insert)
	http.HandleFunc("/users/login", userController.Login)
	http.HandleFunc("/users/detail", userController.GetById)
	fmt.Println("Web Starting")
	http.ListenAndServe(":8085", nil)
}
