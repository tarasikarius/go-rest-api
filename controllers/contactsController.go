package controllers

import (
	"encoding/json"
	"github.com/tarasikarius/go-rest-api/models"
	u "github.com/tarasikarius/go-rest-api/utils"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()

	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user") . (uint)

	data := models.GetContacts(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data

	u.Respond(w, resp)
}
