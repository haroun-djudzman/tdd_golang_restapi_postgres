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
	CreateUserByName(name string)
}

type UserServer struct {
	Retriever UserRetriever
}

func (u *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.getUser(w, r)
	case http.MethodPost:
		u.createUser(w, r)
	}
}

func (u *UserServer) getUser(w http.ResponseWriter, r *http.Request) {
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

func (u *UserServer) createUser(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/")
	u.Retriever.CreateUserByName(name)
	w.WriteHeader(http.StatusAccepted)
}
