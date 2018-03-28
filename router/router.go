package router

import (
	"github.com/buaazp/fasthttprouter"
)

func InitRouter() *fasthttprouter.Router {
	r := fasthttprouter.New()

	r.POST("/forum/:slug", nil)
	r.POST("/forum/:slug/create", nil)
	r.GET("/forum/:slug/detail", nil)
	r.GET("/forum/:slug/threads", nil)
	r.GET("/forum/:slug/users", nil)

	r.GET("/post/:id/details", nil)
	r.POST("/post/:id/details", nil)

	r.POST("/service/clear", nil)
	r.GET("/service/status", nil)

	r.POST("/thread/:slug_or_id/create", nil)
	r.GET("/thread/:slug_or_id/details", nil)
	r.POST("/thread/:slug_or_id/details", nil)
	r.GET("/thread/:slug_or_id/posts", nil)
	r.POST("/thread/:slug_or_id/vote", nil)

	r.POST("/user/:nickname/create", nil)
	r.GET("/user/:nickname/profile", nil)
	r.POST("/user/:nickname/profile", nil)

	return r
}
