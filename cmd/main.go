// cmd/main.go
package main

import (
	"log"

	"github.com/shriansh2002/gowithmysql/internal/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.Run()
}
