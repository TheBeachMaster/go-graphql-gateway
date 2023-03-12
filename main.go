package main

import (
	"log"
	"os"

	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"

	"com.thebeachmaster/nautilusgw/config"
	"com.thebeachmaster/nautilusgw/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("Starting Server...")

	/// Configs
	var cfg *config.Config

	if os.Getenv("USE_FILE_CONFIG") == "true" {

		appConfigPath := "./config/config"

		cfgFile, err := config.LoadConfig(appConfigPath)
		if err != nil {
			log.Fatalf("LoadConfig Error: %s", err.Error())
		}

		pc, err := config.ParseConfig(cfgFile)
		if err != nil {
			log.Fatalf("ParseConfig Error: %s", err.Error())
		}

		cfg = pc
	} else {

		lc, err := config.LoadEnvConfig()
		if err != nil {
			log.Fatalf("ParseConfig Error: %s", err.Error())
		}
		cfg = lc
	}
	/// Gateway Init
	// Load the endpoints here first before anything...
	schemas, err := graphql.IntrospectRemoteSchemas(cfg.GraphQLEndpoints...)
	if err != nil {
		log.Fatalf("Schema loading error: %s", err.Error())
	}

	gw, err := gateway.New(schemas)
	if err != nil {
		log.Fatalf("Gateway initialization error: %s", err.Error())
	}

	/// Server Init
	server := server.NewServer(cfg, gw)
	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}
