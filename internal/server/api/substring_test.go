package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstring(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "abcabcbb",
			input:    "abcabcbb",
			expected: "Max len:3, String:abc",
		},
		{
			name:     "bbbb",
			input:    "bbbb",
			expected: "Max len:1, String:b",
		},
		{
			name:     "pwwkew",
			input:    "pwwkew",
			expected: "Max len:3, String:wke",
		},
		{
			name:     "copy",
			input:    "copy",
			expected: "Max len:4, String:copy",
		},
		{
			name:     "c",
			input:    "c",
			expected: "Max len:1, String:c",
		},
	}

	handler := http.HandlerFunc(Substring)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("api/substring/?str=%s", tc.input), nil)
			handler.ServeHTTP(rec, req)
			assert.Equal(t, tc.expected, rec.Body.String())
		})
	}
}
