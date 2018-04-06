package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
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
	}
}

func PostDetails(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	//id := ctx.UserValue("id").(string)
	//related := ctx.FormValue("related")
	//array := strings.Split(string(related), ",")
	//log.Println(id)
	//check parameters
	//for _, arr := range array {
	//	log.Println(arr)
	//}
}
