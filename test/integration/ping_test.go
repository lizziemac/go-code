package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"myapp/internal/api/server"
)

// TestPingEndpoint verifies that the /api/v1/ping endpoint responds with a 200 OK status.
// Intentionally uses httptest.NewServer to test the full stack including middleware.
func TestPingEndpoint(t *testing.T) {
	srv := httptest.NewServer(server.Setup())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/api/v1/ping")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}

	var body struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if body.Message != "pong" {
		t.Errorf("expected pong, got %q", body.Message)
	}
}
