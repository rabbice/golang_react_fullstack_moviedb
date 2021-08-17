package main

import (
	"log"

	"github.com/rabbice/movieapi/src/backend/routes"
)

func main() {
	log.Println("Main log...")
	log.Fatal(routes.RunAPI(":8000"))
}
