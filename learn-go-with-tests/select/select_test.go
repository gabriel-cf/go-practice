package selector

import (
	"net/http"
	"net/http/httptest"
	"time"
	"testing"
)

func TestRacer(t *testing.T) {
	t.Run("Should return the URL of the fastest server", func(t *testing.T) {
		slowServer := startTestServer(200 * time.Millisecond)
		fastServer := startTestServer(0 * time.Millisecond)

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
            t.Fatalf("did not expect an error but got one %v", err)
        }

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("Should throw an error if the server takes longer that timeout s", func(t *testing.T) {
		slowServer := startTestServer(200 * time.Millisecond)
		fastServer := startTestServer(210 * time.Millisecond)

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 190 * time.Millisecond)

		if err == nil {
			t.Errorf("Expected an error but got none")
		}
	})
	
}

func startTestServer(delayInMs time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayInMs)
		w.WriteHeader(http.StatusOK)
	}))
}