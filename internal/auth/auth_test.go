package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Case 1: valid API key
	headers := http.Header{"Authorization": []string{"ApiKey my-key"}}
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if key != "my-key" {
		t.Errorf("expected 'my-key', got %q", key)
	}

	// Case 2: missing Authorization header
	headers = http.Header{}
	_, err = GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}
