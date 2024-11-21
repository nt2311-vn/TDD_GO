package select_mod

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of fastest one", func(t *testing.T) {
		slowServer := makeDelayServer(2 * time.Second)
		fastServer := makeDelayServer(0 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != fastURL {
			t.Errorf("got %q but want %q", got, fastURL)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayServer(25 * time.Second)

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Second)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
