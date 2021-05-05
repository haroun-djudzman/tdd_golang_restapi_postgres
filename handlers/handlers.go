package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserRetriever interface {
	GetUserName(id int) string
}

type UserServer struct {
	Retriever UserRetriever
}

func (u *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/user/"))
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	name := u.Retriever.GetUserName(id)

	if name == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, name)
}
