package main

import (
	"log"
	"net/http"

	"github.com/haroun-djudzman/restapi-postgres/handlers"
)

func main() {
	server := &handlers.UserServer{}
	log.Fatal(http.ListenAndServe(":8081", server))
}
