package app

import (
	"github.com/vers-dev/go-skeleton/config"
	"github.com/vers-dev/go-skeleton/database"
	"github.com/vers-dev/go-skeleton/server"
)

func Run() {
	conf := config.GetConfig()
	database.NewPostgresDatabase(conf)
	server.StartServer()
}
