package controller

import (
	"encoding/json"
	"net/http"

	"github.com/liquidslr/issuetracker/domain"
)

type UserController struct {
	UserService domain.UserService
}

func (c UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserService.Users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJSON, err := json.Marshal(users)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(userJSON)

}

func (c UserController) Show(w http.ResponseWriter, r *http.Request) {

}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c UserController) Delete(w http.ResponseWriter, r *http.Request) {

}
