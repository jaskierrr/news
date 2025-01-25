package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"main/config"
	"os"

	_ "github.com/lib/pq"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects"
)

// const dbConfigString = "postgres://%s:%s@%s:%s/%s"


func NewDB(config config.Config, logger *slog.Logger) *reform.DB {
	// connString := fmt.Sprintf(dbConfigString, config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name)

	dbPostgresConn, err := sql.Open("postgres", dsn)
	if err != nil {
		// TODO fix this fatal
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = dbPostgresConn.Ping()
	if err != nil {
		// TODO fix this fatal
		log.Fatalf("Unable to PING to database: %v\n", err)
	}

	loggerForSQL := log.New(os.Stderr, "SQL: ", log.Flags())

	dbPostgres := reform.NewDB(dbPostgresConn, dialects.ForDriver("postgres"), reform.NewPrintfLogger(loggerForSQL.Printf))

	logger.Info("DB",
		slog.Any("data", dbPostgres))


	return dbPostgres
}
