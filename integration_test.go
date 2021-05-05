package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/haroun-djudzman/restapi-postgres/handlers"
	"github.com/haroun-djudzman/restapi-postgres/testingUtil"
)

func TestCreatingAndGettingUser(t *testing.T) {
	store := NewInMemoryUserStore()
	server := handlers.UserServer{Retriever: store}
	names := []string{"Ilham", "Putri", "Reno"}

	// Create three users
	for _, name := range names {
		server.ServeHTTP(httptest.NewRecorder(), testingUtil.NewCreateUserRequest(name))
	}

	// Get three newly created users
	for i, name := range names {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, testingUtil.NewGetUserRequest(i+1))
		testingUtil.AssertStatus(t, response.Code, http.StatusOK)
		testingUtil.AssertResponseBody(t, response.Body.String(), name)
	}
}
