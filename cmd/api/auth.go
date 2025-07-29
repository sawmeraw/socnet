package main

import (
	"net/http"

	"github.com/sawmeraw/gogo/internal/store"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

// RegisterUser godoc
//
//	@Summary		Registers a user
//	@Description	Registers a user
//	@Tags			authentication
//	@Accept			json
//	@Product		json
//	@Param			payload	bod		RegisterUserPayload	true	"User credentials"
//	@Success		200	{object}	store.User "User Registered"
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router			/auth/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	payload := &RegisterUserPayload{}
	err := readJSON(w, r, payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//hash the user password
	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	if err := user.Password.Set(payload.Password); err != nil {
		app.statusInternalServerError(w, r, err)
		return
	}

	// ctx := r.Context()
	if err := app.jsonResponse(w, http.StatusCreated, nil); err != nil {
		app.statusInternalServerError(w, r, err)
	}

}
