package main

import (
	"net/http"
	"time"

	"github.com/sawmeraw/gogo/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	weekAgo := time.Now().AddDate(0, 0, -7)

	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
		Search: "",
		Tags:   []string{},
		Since:  weekAgo.Format(time.DateTime),
		Until:  time.Now().Format(time.DateTime),
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	posts, err := app.store.Posts.GetUserFeed(ctx, int64(9), fq)

	if err != nil {
		app.statusInternalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, posts); err != nil {
		app.statusInternalServerError(w, r, err)
	}

}
