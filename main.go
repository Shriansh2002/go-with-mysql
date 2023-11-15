package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DBConfig is a struct to store database configuration
type DBConfig struct {
	Username   string
	Password   string
	Host       string
	Port       string
	Database   string
}

func loadDBConfig() (*DBConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %s", err)
	}

	dbConfig := &DBConfig{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Database: os.Getenv("DATABASE"),
	}

	return dbConfig, nil
}

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
	dbConfig, err := loadDBConfig()
	if err != nil {
		fmt.Println("Error:", err)
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