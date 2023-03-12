package graphql

import (
	"net/http"

	"github.com/nautilus/gateway"
)

type grapqhqlHandler struct {
	gw *gateway.Gateway
}

func NewGrapqhQLHandler(gatewayInst *gateway.Gateway) GraphqlRouter {
	return &grapqhqlHandler{gw: gatewayInst}
}

func (g *grapqhqlHandler) Playground() http.HandlerFunc {
	return g.gw.PlaygroundHandler
}

func (g *grapqhqlHandler) Service() http.HandlerFunc {
	return g.gw.GraphQLHandler
}
