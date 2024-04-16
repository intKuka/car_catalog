package initializers

import (
	"car_catalog/internal/config"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

// TODO: add table constraints
func OpenConnection() {
	var err error
	DB, err = pgx.Connect(context.Background(), config.Cfg.ConnectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// TODO: set required fields
	sql := `
		CREATE TABLE IF NOT EXISTS Cars (
			Id SERIAL PRIMARY KEY,
			RegNum CHARACTER VARYING,
			Mark CHARACTER VARYING,
			Model CHARACTER VARYING,
			Year INTEGER,
			Name CHARACTER VARYING(30),
			Surname CHARACTER VARYING(30),
			Patronymic CHARACTER VARYING(30)
		);`

	_, err = DB.Exec(context.Background(), sql)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to prepare table: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Conncected to DB")
}
