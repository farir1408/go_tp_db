package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
	"log"
	"strings"
)

func PostsCreate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug_or_id").(string)
	posts := models.Posts{}
	posts.UnmarshalJSON(ctx.PostBody())

	err := posts.PostsCreate(slug)

	if err == errors.ThreadNotFound {
		resErr := models.Error{}
		ctx.SetStatusCode(404)
		log.Println("this block is completed ThreadNotFound Posts")
		ctx.Write(resErr.ErrorMsgJSON(err.Error()))
	}
	//TODO: err == nil, it's necessary to finish
}

func PostDetails(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	id := ctx.UserValue("id").(string)
	related := ctx.FormValue("related")
	array := strings.Split(string(related), ",")
	log.Println(id)
	//check parameters
	for _, arr := range array {
		log.Println(arr)
	}
	//TODO: important
}
