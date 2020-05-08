package controllers

import (
	"encoding/json"
	"github.com/tarasikarius/go-rest-api/models"
	u "github.com/tarasikarius/go-rest-api/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request data"))

		return
	}

	resp := account.Create()

	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))

		return
	}

	resp := models.Login(account.Email, account.Password)

	u.Respond(w, resp)
}
