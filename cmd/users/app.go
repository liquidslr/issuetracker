package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/liquidslr/issuetracker/application"
	"github.com/liquidslr/issuetracker/domain"
	"github.com/liquidslr/issuetracker/web/controller"

	memory "github.com/liquidslr/issuetracker/persistence/db"
)

func main() {
	userRepo := memory.NewUserRepository()

	userService := application.UserService{
		UsersRepository: userRepo,
	}

	userController := controller.UserController{
		UserService: userService,
	}

	for i := 0; i < 10; i++ {
		userService.Create(&domain.User{Name: fmt.Sprintf("User_%d", i)})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", userController.List)

	server := &http.Server{
		Addr:           ":8090",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
