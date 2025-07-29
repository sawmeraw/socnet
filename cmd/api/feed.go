package main

import (
	"net/http"
	"time"

	"github.com/sawmeraw/gogo/internal/store"
)

// GetUserFeed godoc
//
//	@Summary		Fetch user feed
//	@Description	Fetches the user feed by userID and params
//	@Tags			feed
//	@Accept			json
//	@Product		json
//	@Param			since	query		string	false	"Since"
//	@Param			until	query		string	false	"Until"
//	@Param			limit	query		int		false	"Limit"
//	@Param			offset	query		int		false	"Offset"
//	@Param			search	query		string	false	"Search"
//	@Param			sort	query		string	false	"Sort"
//	@Param			tags	query		string	false	"Tags"
//	@Success		200		{object}	[]store.PostWithMetadata
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/feed [get]
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
