package server

import (
	"log"
	"os"
	"os/signal"
	"time"

	"com.thebeachmaster/nautilusgw/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nautilus/gateway"
)

type Server struct {
	app     *fiber.App
	cfg     *config.Config
	gateway *gateway.Gateway
}

var Version = "1.0.0"

func NewServer(cfg *config.Config, gw *gateway.Gateway) *Server {
	_app := fiber.New(fiber.Config{
		Prefork:      cfg.Server.Prefork,
		ReadTimeout:  time.Second * time.Duration(cfg.Server.ReadTimeout),
		AppName:      cfg.Server.AppName + " Version: " + Version,
		ServerHeader: cfg.Server.ServerHeader,
	})

	// You can add app-level middlewares here
	_app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	return &Server{app: _app, cfg: cfg, gateway: gw}
}

func (srv *Server) Run() error {
	go func() {
		log.Printf("Server is listening on PORT: %s", srv.cfg.Port)
		addr := ":" + srv.cfg.Port
		if err := srv.app.Listen(addr); err != nil {
			log.Panicf("[CRIT] Unable to start server. Reason: %v", err)
		}
	}()

	quitServer := make(chan struct{})

	err := srv.MapHandlers(srv.app)
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	close(quitServer)

	<-quitServer

	log.Printf("Server shutdown")
	return srv.app.Shutdown()

}
