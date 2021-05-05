package main

import (
	"log"
	"net/http"

	"github.com/haroun-djudzman/restapi-postgres/handlers"
)

type InMemoryUserRetriever struct {
}

func (i *InMemoryUserRetriever) GetUserName(id int) string {
	return "John Doe"
}

func main() {
	server := &handlers.UserServer{Retriever: &InMemoryUserRetriever{}}
	log.Fatal(http.ListenAndServe(":8081", server))
}
