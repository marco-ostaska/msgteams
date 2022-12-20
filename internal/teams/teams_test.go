package teams_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marco-ostaska/msgteams/internal/teams"
)

func TestNewPostMsg(t *testing.T) {
	// Create a test server that will return a 200 status code.
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Set the test server URL as the target for the function.
	url := server.URL

	tt := []struct {
		name string
		url  string
		msg  string
		err  error
	}{
		{"Post sent", url, "hello world", fmt.Errorf("Test OK")},
		{"Wrong Url", "https://nada", "hello world", nil},
		{"NO URL", "", "hello world", nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function with the test server URL and a message.
			if err := teams.NewPostMsg(tc.url, tc.msg); err == tc.err {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}

}
