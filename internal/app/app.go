package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shriansh2002/gowithmysql/pkg/config"
	"github.com/shriansh2002/gowithmysql/pkg/database"
)

// Run initializes and runs the application.
func Run() {
	dbConfig, err := config.LoadDBConfig()
	if err != nil {
		log.Fatalf("Error loading database configuration: %s", err)
		return
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)

	tableName := "Stack"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %s", err)
	}

	// database.CreateTable(db, tableName)
	// database.InsertData(db, tableName)
	database.ShowData(db, tableName)
	// database.DropTable(db, tableName)
}
