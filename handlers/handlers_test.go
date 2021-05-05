package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/haroun-djudzman/restapi-postgres/handlers"
)

type StubUserRetriever struct {
	names map[int]string
}

func (s *StubUserRetriever) GetUserName(id int) string {
	name := s.names[id]
	return name
}

func TestGetUser(t *testing.T) {
	retriever := StubUserRetriever{
		map[int]string{
			1: "Budi",
			2: "Siti",
		},
	}
	server := &handlers.UserServer{&retriever}

	t.Run("get budi name by id", func(t *testing.T) {
		request := newGetUserRequest(1)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Budi"
		assertResponseBody(t, got, want)
	})

	t.Run("get siti name by id", func(t *testing.T) {
		request := newGetUserRequest(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Siti"
		assertResponseBody(t, got, want)
	})
}

func newGetUserRequest(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/user/%d", id), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
