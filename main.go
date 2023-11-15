package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func showData(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("SELECT * FROM %s", tableName)

	rows, err := db.Query(sqlCommand)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var isPublished bool
		err = rows.Scan(&id, &name, &isPublished)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name, isPublished)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}

func createTable(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("CREATE TABLE %s (id int PRIMARY KEY, name varchar(255), isPublished boolean)", tableName)

	_, err := db.Exec(sqlCommand)
	if err != nil {
		panic(err.Error())
	}
}

func dropTable(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("DROP TABLE %s", tableName)

	_, err := db.Exec(sqlCommand)
	if err != nil {
		panic(err.Error())
	}
}

func insertData(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("INSERT INTO %s(id, name, isPublished) VALUES(?, ?, ?)", tableName)

	stmt, err := db.Prepare(sqlCommand)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(1, "Golang", true)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(2, "Python", true)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(3, "Java", true)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil{
	log.Fatalf("Error loading .env file: %s", err)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	tableName := "Stack"

    db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// createTable(db, tableName)
	// insertData(db, tableName)
	showData(db, tableName)
	// dropTable(db, tableName)
}