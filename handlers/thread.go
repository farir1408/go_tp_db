package handlers

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
)

func ThreadDetails(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug_or_id").(string)
	thread := models.Thread{}

	err := thread.ThreadDetails(slug)

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := thread.MarshalJSON()
		ctx.Write(buf)
	case errors.ThreadNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	}
}

func ThreadUpdateDetails(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug_or_id").(string)
	thread := models.ThreadUpdate{}
	thread.UnmarshalJSON(ctx.PostBody())

	result, err := thread.ThreadUpdate(slug)

	if err == nil {
		ctx.SetStatusCode(200)
		buf, _ := result.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.ThreadNotFound {
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	}
}

func Vote(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	vote := models.Vote{}
	vote.UnmarshalJSON(ctx.PostBody())
	slug := ctx.UserValue("slug_or_id").(string)

	err := vote.Vote(slug)

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		thread := models.Thread{}
		thread.ThreadDetails(slug)
		buf, _ := thread.MarshalJSON()
		ctx.Write(buf)
	case errors.ThreadNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	}
}

func ThreadPosts(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug_or_id").(string)
	limit := ctx.FormValue("limit")
	since := ctx.FormValue("since")
	sort := ctx.FormValue("sort")
	desc := ctx.FormValue("desc")

	threadId, err := models.GetPostThreadId(slug)

	var posts *models.Posts
	if err == nil {
		switch true {
		case bytes.Equal(sort, []byte("tree")):
			posts, err = models.GetPostsSortTree(threadId, limit, since, desc)
			break

		case bytes.Equal(sort, []byte("parent_tree")):
			posts, err = models.GetPostsSortParentTree(threadId, limit, since, desc)
			break

		default:
			posts, err = models.GetPostsSortFlat(threadId, limit, since, desc)
			break
		}
	}

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := posts.MarshalJSON()
		ctx.Write(buf)
	case errors.ThreadNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	}
}
