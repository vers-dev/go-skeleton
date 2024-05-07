package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/vers-dev/go-skeleton/config"
	"sync"
)

type postgresDatabase struct {
	Db *pgx.Conn
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			conf.Db.Host,
			conf.Db.User,
			conf.Db.Password,
			conf.Db.DBName,
			conf.Db.Port,
			conf.Db.SSLMode,
			conf.Db.TimeZone,
		)

		db, err := pgx.Connect(context.Background(), dsn)

		if err != nil {
			panic("Error db connection")
		}

		dbInstance = &postgresDatabase{Db: db}
	})

	err := dbInstance.Db.Ping(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Println("Database successful connected!")

	return dbInstance
}

func (p *postgresDatabase) GetDb() *pgx.Conn {
	return dbInstance.Db
}
