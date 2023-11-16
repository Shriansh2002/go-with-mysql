// pkg/database/connection.go

package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/shriansh2002/gowithmysql/pkg/config"
)

// ConnectDB connects to the database and returns a pointer to the sql.DB object.
func ConnectDB() (*sql.DB, error){
	dbConfig, err := config.LoadDBConfig()
	if err != nil {
		log.Fatalf("Error loading database configuration: %s", err)
		return nil, err
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)


	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %s", err)
		return nil, err
	}

	return db, nil
}