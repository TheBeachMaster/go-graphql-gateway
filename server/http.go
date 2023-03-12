package server

import (
	"time"

	"com.thebeachmaster/nautilusgw/graphql"
	"github.com/gofiber/fiber/v2"
)

func (srv *Server) MapHandlers(app *fiber.App) error {

	app.Get("/ruok", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "imok",
			"time":   time.Now().String(),
		})
	})

	graphqlRoute := app.Group("/api")

	handler := graphql.NewGrapqhQLHandler(srv.gateway)
	graphql.MapGraphqlRouter(graphqlRoute, handler)

	return nil
}
