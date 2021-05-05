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
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	fmt.Fprint(w, u.Retriever.GetUserName(id))
}

// func getUserName(id int) string {
// 	if id == 1 {
// 		return "Budi"
// 	}

// 	if id == 2 {
// 		return "Siti"
// 	}
// 	return ""
// }
