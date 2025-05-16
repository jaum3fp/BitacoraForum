package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/controllers"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

func setPostRoutes(router fiber.Router, repo repositorys.PostRepository) {

	postHandler := controllers.NewPostHandler(repo)

	group := router.Group("/post")

	group.Get("/all", postHandler.GetAllPosts)
	group.Post("/", postHandler.CreatePost)

	group.Get("/:id", postHandler.GetPost)
	group.Put("/:id", postHandler.UpdatePost)
	group.Delete("/:id", postHandler.DeletePost)

	group.Get("/tag/:id", postHandler.GetPostsByTag)
	group.Patch("/inc-view/:id", postHandler.IncrementPostViews)

	group.Get("/count/:flag", postHandler.GetCountryPostsNumber)
	group.Get("/all/:cc", postHandler.GetCountryPosts)
}
