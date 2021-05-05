package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/haroun-djudzman/restapi-postgres/handlers"
	"github.com/haroun-djudzman/restapi-postgres/testingUtil"
)

type StubUserStore struct {
	names map[int]string
}

func (s *StubUserStore) GetUserName(id int) string {
	name := s.names[id]
	return name
}

func (s *StubUserStore) CreateUserByName(name string) {
	id := len(s.names) + 1
	s.names[id] = name
}

func TestGetUser(t *testing.T) {
	store := StubUserStore{
		map[int]string{
			1: "Budi",
			2: "Siti",
		},
	}
	server := &handlers.UserServer{&store}

	t.Run("get budi name by id", func(t *testing.T) {
		request := testingUtil.NewGetUserRequest(1)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		testingUtil.AssertStatus(t, response.Code, http.StatusOK)
		got := response.Body.String()
		want := "Budi"
		testingUtil.AssertResponseBody(t, got, want)
	})

	t.Run("get siti name by id", func(t *testing.T) {
		request := testingUtil.NewGetUserRequest(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		testingUtil.AssertStatus(t, response.Code, http.StatusOK)
		got := response.Body.String()
		want := "Siti"
		testingUtil.AssertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing user", func(t *testing.T) {
		request := testingUtil.NewGetUserRequest(3)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		testingUtil.AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestCreateUser(t *testing.T) {
	store := StubUserStore{
		map[int]string{},
	}
	server := &handlers.UserServer{&store}

	t.Run("create new user on POST", func(t *testing.T) {
		name := "Anto"
		request := testingUtil.NewCreateUserRequest(name)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		testingUtil.AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.names) != 1 {
			t.Errorf("no new user is inserted, %d user in database", len(store.names))
		}

		if store.names[1] != name {
			t.Errorf("created wrong user, got %q want %q", store.names[1], name)
		}
	})
}
