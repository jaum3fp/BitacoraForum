package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/controllers"
	"github.com/jaum3fp/bitacora-forum/internal/middleware"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

func setPostRoutes(router fiber.Router, repo repositorys.PostRepository) {

	postHandler := controllers.NewPostHandler(repo)

	group := router.Group("/post")

	group.Get("/all", postHandler.GetAllPosts)

	group.Post("/", postHandler.CreatePost)
	group.Get("/:id", postHandler.GetPost)
	group.Put("/:id", middleware.Protected(), postHandler.UpdatePost)
	group.Delete("/:id", middleware.Protected(), postHandler.DeletePost)

	group.Get("/tag/:id", postHandler.GetPostsByTag)
	group.Patch("/inc-view/:id", postHandler.IncrementPostViews)
	group.Get("/comments/:super", postHandler.GetPostComments)
	group.Get("/count/comments/:super", postHandler.GetPostCommentsNumber)
	group.Get("/count/:flag", postHandler.GetCountryPostsNumber)

}
