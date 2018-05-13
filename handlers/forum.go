package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
	"strings"
)

func ForumCreate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	forum := models.Forum{}
	forum.UnmarshalJSON(ctx.PostBody())

	resp, err := forum.ForumCreate()

	switch err {
	case nil:
		ctx.SetStatusCode(201)
		buf, _ := forum.MarshalJSON()
		ctx.Write(buf)
	case errors.UserNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	case errors.ForumIsExist:
		ctx.SetStatusCode(409)
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	}
}

func ForumDetails(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug").(string)
	strings.ToLower(slug)
	forum := models.Forum{}
	err := forum.ForumDetails(slug)

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := forum.MarshalJSON()
		ctx.Write(buf)
	case errors.UserNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	}
}

func ForumThreadCreate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	thread := models.Thread{}
	thread.UnmarshalJSON(ctx.PostBody())

	slug := ctx.UserValue("slug").(string)
	thread.ForumId = slug

	resp, err := thread.ForumThreadCreate()

	switch err {
	case nil:
		ctx.SetStatusCode(201)
		buf, _ := thread.MarshalJSON()
		ctx.Write(buf)
	case errors.ForumNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	case errors.ThreadIsExist:
		ctx.SetStatusCode(409)
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	}
}

func GetThreads(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug").(string)
	limit := ctx.FormValue("limit")
	since := ctx.FormValue("since")
	desc := ctx.FormValue("desc")
	resp, err := models.GetThreads(slug, limit, since, desc)

	switch err {
	case errors.ForumNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	}
}

func GetUsers(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug").(string)
	limit := ctx.FormValue("limit")
	since := ctx.FormValue("since")
	desc := ctx.FormValue("desc")

	resp, err := models.GetUsers(slug, limit, since, desc)

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	case errors.ForumNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	}
}
