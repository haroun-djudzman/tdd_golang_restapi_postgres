package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/haroun-djudzman/restapi-postgres/handlers"
)

func TestGetUser(t *testing.T) {
	t.Run("get budi name by id", func(t *testing.T) {
		request := newGetUserRequest(1)
		response := httptest.NewRecorder()

		handlers.UserServer(response, request)

		got := response.Body.String()
		want := "Budi"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("get siti name by id", func(t *testing.T) {
		request := newGetUserRequest(2)
		response := httptest.NewRecorder()

		handlers.UserServer(response, request)

		got := response.Body.String()
		want := "Siti"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
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
