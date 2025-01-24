package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers        http.Header // Input headers
		expectedAPIKey string      // Expected API key
		expectedError  error       // Expected error
	}{
		"valid header": {
			headers:        http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			expectedAPIKey: "my-secret-key",
			expectedError:  nil,
		},
		"missing authorization header": {
			headers:        http.Header{},
			expectedAPIKey: "",
			expectedError:  ErrNoAuthHeaderIncluded,
		},
		"empty authorization header": {
			headers:        http.Header{"Authorization": []string{""}},
			expectedAPIKey: "",
			expectedError:  ErrNoAuthHeaderIncluded,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)

			// Compare API Key
			if apiKey != tc.expectedAPIKey {
				t.Errorf("expected API key %q, got %q", tc.expectedAPIKey, apiKey)
			}

			// Compare Errors
			if (err != nil && tc.expectedError == nil) || (err == nil && tc.expectedError != nil) || (err != nil && err.Error() != tc.expectedError.Error()) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
