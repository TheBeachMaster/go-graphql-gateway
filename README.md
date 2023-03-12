# Go GraphQL Gateway

- A dead-simple GraphQL gateway for all your federated Go GraphQL server needs.

Uses [Nautilus](https://github.com/nautilus/gateway) and [Fiber](https://github.com/gofiber/fiber)

## TLDR;

If you're just looking for a quick and dirty implementation of a federated gateway, paste this code into your `main.go` file, package and ship.

```go 
package main

import (
	"fmt"
	"net/http"

	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
)

func main() {
	// change this to point to your GraphQL servers
	schemas, err := graphql.IntrospectRemoteSchemas(
		"http://localhost:4000/graphql",
		"http://localhost:4001/graphql",
        "http://as.many.urls.as.you.want:6969/graphql"
	)
	if err != nil {
		panic(err)
	}

	// create the gateway instance
	gw, err := gateway.New(schemas)
	if err != nil {
		panic(err)
	}

    // your GraphQL endpoint is http(s)://host(:port)/graphql <- executes queries on POST
    // navigating to this url on the browser will show the GraphQL playground UI
	http.HandleFunc("/graphql", gw.PlaygroundHandler)

	// start the server
	fmt.Println("Starting server")
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
```


## Usage

- Create a new `.env` file based on `.env.example` file
	- `GRAPHQL_ENDPOINTS` variable should be in the format `"http://yourendpoint1","http://youendpoint2",...`
	- `USE_FILE_CONFIG` variable should be set to `true` if you want to set/use your configs inside `config/config.yaml` file (useful for K8s-like deployments).

- Your app will be running on port `8082` and your endpoint will be `api/graphql` and `api/playground` for the playground

- Build the binary `make build` or `make build-mac` for Mac OS

 > Alternatively you can build a Docker image with the Dockerfile provided