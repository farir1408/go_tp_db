package router

import (
	"github.com/buaazp/fasthttprouter"
	"go_tp_db/handlers"
)

func InitRouter() *fasthttprouter.Router {
	r := fasthttprouter.New()

	r.POST("/forum/:slug", handlers.ForumCreate)
	r.POST("/forum/:slug/create", handlers.ForumThreadCreate)
	r.GET("/forum/:slug/details", handlers.ForumDetails)
	r.GET("/forum/:slug/threads", handlers.GetThreads)
	r.GET("/forum/:slug/users", nil)

	r.GET("/post/:id/details", handlers.PostDetails) //TODO: important
	r.POST("/post/:id/details", nil)

	r.POST("/service/clear", nil)
	r.GET("/service/status", nil)

	r.POST("/thread/:slug_or_id/create", handlers.PostsCreate) //TODO: important
	r.GET("/thread/:slug_or_id/details", handlers.ThreadDetails)
	r.POST("/thread/:slug_or_id/details", handlers.ThreadUpdateDetails)
	r.GET("/thread/:slug_or_id/posts", nil)
	r.POST("/thread/:slug_or_id/vote", handlers.Vote)

	r.POST("/user/:nickname/create", handlers.UserCreate)
	r.GET("/user/:nickname/profile", handlers.UserProfile)
	r.POST("/user/:nickname/profile", handlers.UserUpdateProfile)

	return r
}
