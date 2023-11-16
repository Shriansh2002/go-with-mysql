// pkg/database/database.go

package database

import (
	"database/sql"
	"fmt"
	"log"
)

// ShowData fetches and displays data from the specified table.
func ShowData(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("SELECT * FROM %s", tableName)

	rows, err := db.Query(sqlCommand)
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var isPublished bool
		err = rows.Scan(&id, &name, &isPublished)
		if err != nil {
			log.Fatal("Error fetching data:", err)
		}
		fmt.Println(id, name, isPublished)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error fetching data:", err)
	}
}

// CreateTable creates a table in the database.
func CreateTable(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("CREATE TABLE %s (id int PRIMARY KEY, name varchar(255), isPublished boolean)", tableName)

	_, err := db.Exec(sqlCommand)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

// DropTable drops a table from the database.
func DropTable(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("DROP TABLE %s", tableName)

	_, err := db.Exec(sqlCommand)
	if err != nil {
		log.Fatal("Error dropping table:", err)
	}
}

// InsertData inserts sample data into the specified table.
func InsertData(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("INSERT INTO %s(id, name, isPublished) VALUES(?, ?, ?)", tableName)

	stmt, err := db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal("Error in preparing SQL statement:", err)
	}

	_, err = stmt.Exec(1, "Golang", true)
	if err != nil {
		log.Fatal("Error in Inserting Data", err)
	}

	_, err = stmt.Exec(2, "Python", true)
	if err != nil {
		log.Fatal("Error in Inserting Data", err)
	}

	_, err = stmt.Exec(3, "Java", true)
	if err != nil {
		log.Fatal("Error in Inserting Data", err)
	}
}
