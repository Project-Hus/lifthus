package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"routine/ent"
)

// ConncectToLifthusRoutine returns lifthus_routine_db's ent client.
// you've got to close it with Close() in defer out of this function.
func ConnectToLifthusRoutine() (*ent.Client, error) {
	dbHost, ok1 := os.LookupEnv("LIFTHUS_ROUTINE_DB_HOST")
	dbPort, ok2 := os.LookupEnv("LIFTHUS_ROUTINE_DB_PORT")
	dbUser, ok3 := os.LookupEnv("LIFTHUS_ROUTINE_DB_USER")
	dbPassword, ok4 := os.LookupEnv("LIFTHUS_ROUTINE_DB_PASSWORD")
	dbName, ok5 := os.LookupEnv("LIFTHUS_ROUTINE_DB_NAME")
	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 {
		log.Fatal("LIFTHUS_ROUTINE_DB credentials are not set")
	}

	// DB connection
	connectionPhrase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := ent.Open("mysql", connectionPhrase)
	if err != nil {
		log.Print("[F] opening connection to mysql failed: %w", err)
		return nil, err
	}

	// Running the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
	); err != nil {
		log.Print(" creating schema resources failed: %w", err)
		return nil, err
	}

	return client, nil
}
