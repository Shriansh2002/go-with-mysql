// internal/app/main.go

package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shriansh2002/gowithmysql/pkg/database"
)

// Run initializes and runs the application.
func Run() {
	tableName := "Stack"

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// database.CreateTable(db, tableName)
	// database.InsertData(db, tableName)
	database.ShowData(db, tableName)
	// database.DropTable(db, tableName)
}
