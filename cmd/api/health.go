package main

import (
	"net/http"
)

// AppHealth godoc
//
//	@Summary		Health check
//	@Description	Responds with health details about the server
//	@Tags			health
//	@Accept			json
//	@Product		json
//	@Success		200	{string}	string	"Server is up"
//	@Failure		500	{object}	error	"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.statusInternalServerError(w, r, err)
	}
}
