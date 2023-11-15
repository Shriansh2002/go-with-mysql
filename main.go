package main

import (
	"database/sql"
	"fmt"
	"log"
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

func createTable(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("CREATE TABLE %s (id int PRIMARY KEY, name varchar(255), isPublished boolean)", tableName)

	_, err := db.Exec(sqlCommand)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

func dropTable(db *sql.DB, tableName string) {
	sqlCommand := fmt.Sprintf("DROP TABLE %s", tableName)

	_, err := db.Exec(sqlCommand)
	if err != nil {
		log.Fatal("Error dropping table:", err)
	}
}

func insertData(db *sql.DB, tableName string) {
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

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dbConfig, err := loadDBConfig()
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

	// createTable(db, tableName)
	// insertData(db, tableName)
	showData(db, tableName)
	// dropTable(db, tableName)
}