package graphql

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type GraphqlRouter interface {
	Playground() http.HandlerFunc
	Service() http.HandlerFunc
}

func MapGraphqlRouter(f fiber.Router, r GraphqlRouter) {
	f.All("/playground", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(r.Playground())(c.Context())
		return nil
	})

	f.Post("/graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(r.Service())(c.Context())
		return nil
	})

}
