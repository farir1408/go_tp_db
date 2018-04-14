package router

import (
	"github.com/buaazp/fasthttprouter"
	"go_tp_db/handlers"
)

func InitRouter() *fasthttprouter.Router {
	r := fasthttprouter.New()

	r.POST("/api/forum/:slug", handlers.ForumCreate)
	r.POST("/api/forum/:slug/create", handlers.ForumThreadCreate)
	r.GET("/api/forum/:slug/details", handlers.ForumDetails)
	r.GET("/api/forum/:slug/threads", handlers.GetThreads)
	r.GET("/api/forum/:slug/users", handlers.GetUsers)

	r.GET("/api/post/:id/details", handlers.PostDetails)
	r.POST("/api/post/:id/details", handlers.PostUpdate)

	r.POST("/api/service/clear", handlers.ClearDataBase)
	r.GET("/api/service/status", handlers.StatusDataBase)

	r.POST("/api/thread/:slug_or_id/create", handlers.PostsCreate)
	r.GET("/api/thread/:slug_or_id/details", handlers.ThreadDetails)
	r.POST("/api/thread/:slug_or_id/details", handlers.ThreadUpdateDetails)
	r.GET("/api/thread/:slug_or_id/posts", handlers.ThreadPosts)
	r.POST("/api/thread/:slug_or_id/vote", handlers.Vote)

	r.POST("/api/user/:nickname/create", handlers.UserCreate)
	r.GET("/api/user/:nickname/profile", handlers.UserProfile)
	r.POST("/api/user/:nickname/profile", handlers.UserUpdateProfile)

	return r
}
