package app

import (
	"go-database/config"
	"go-database/database"
	"go-database/server"
)

func Run() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
