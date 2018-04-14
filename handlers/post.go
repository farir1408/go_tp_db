package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
	"strings"
)

func PostsCreate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug_or_id").(string)
	posts := models.Posts{}
	posts.UnmarshalJSON(ctx.PostBody())

	err := posts.PostsCreate(slug)

	switch err {
	case nil:
		ctx.SetStatusCode(201)
		buf, _ := posts.MarshalJSON()
		ctx.Write(buf)
	case errors.NoPostsForCreate:
		ctx.SetStatusCode(201)
		ctx.Write(ctx.PostBody())
	case errors.ThreadNotFound:
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		ctx.SetStatusCode(404)
		ctx.Write(resErr)
	case errors.NoThreadParent:
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		ctx.SetStatusCode(409)
		ctx.Write(resErr)
	}
}

func PostDetails(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	id := ctx.UserValue("id").(string)
	related := ctx.FormValue("related")
	array := strings.Split(string(related), ",")

	//check parameters
	resp, err := models.PostDetails(id, array)

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	case errors.PostNotFound:
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		ctx.SetStatusCode(404)
		ctx.Write(resErr)
	}
}

func PostUpdate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	id := ctx.UserValue("id").(string)
	update := models.PostUpdate{}
	update.UnmarshalJSON(ctx.PostBody())
	post := models.Post{}

	err := post.PostUpdate(&update, id)
	switch err {
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := post.MarshalJSON()
		ctx.Write(buf)
	case errors.PostNotFound:
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		ctx.SetStatusCode(404)
		ctx.Write(resErr)
	}

}
