package download

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"net/http/httptest"
)

func TestContentBytes(t *testing.T) {
	t.Run("should return no errors", func(t *testing.T) {
		// Create a test server with a mock handler
		mockResponse := "This is a test response"

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		// Call ContentBytes with the test server's URL
		bytes, err := ContentBytes(ts.URL)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, []byte(mockResponse), bytes)
	})

	t.Run("should return error", func(t *testing.T) {
		t.Run("with invalid URL", func(t *testing.T) {
			// Test with an invalid URL
			invalidURL := "http://invalid-url"

			bytes, err := ContentBytes(invalidURL)

			// Assertions
			assert.Error(t, err)
			assert.Empty(t, bytes)
			assert.Contains(t, err.Error(), "error making HTTP request to")
		})

		t.Run("with 404 status error", func(t *testing.T) {
			// Create a test server that returns a 404 status
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			}))
			defer ts.Close()

			// Call ContentBytes with the test server's URL
			bytes, err := ContentBytes(ts.URL)

			// Assertions
			assert.Error(t, err)
			assert.Empty(t, bytes)
			assert.Contains(t, err.Error(), "error response status")
		})
	})
}
