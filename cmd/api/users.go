package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sawmeraw/gogo/internal/store"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	user, err := app.store.Users.GetByID(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.badRequestResponse(w, r, err)

		default:
			app.statusInternalServerError(w, r, err)

		}
		return

	}
	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.statusInternalServerError(w, r, err)
	}

}
