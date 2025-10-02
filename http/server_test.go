package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIClient_GetData(t *testing.T) {
	tests := []struct {
		name           string
		serverResponse string
		serverStatus   int
		expectedMsg    string
		expectedErrMsg string
		expectedCode   int
	}{
		{
			name:           "Successful response",
			serverResponse: `{"message": "success"}`,
			serverStatus:   http.StatusOK,
			expectedMsg:    "success",
		},
		{
			name:           "Not Found response",
			serverResponse: `{"error": "not found"}`,
			serverStatus:   http.StatusNotFound,
			expectedErrMsg: "not found",
		},
		{
			name:           "Internal Server Error",
			serverResponse: `{"error": "internal error"}`,
			serverStatus:   http.StatusInternalServerError,
			expectedErrMsg: "internal error",
		},
	}

	for _, test := range tests {
		// create a new HTTP test server for each test case
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(test.serverStatus)
			w.Write([]byte(test.serverResponse))
		}))
		defer server.Close()

		client := &APIClient{BaseURL: server.URL}

		t.Run(test.name, func(t *testing.T) {
			resp, err := client.GetData()

			if resp == nil {
				t.Fatal("Expected response, got nil")
			} else if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			// check the response message or error message
			if test.expectedMsg != "" && resp.Message != test.expectedMsg {
				t.Errorf("Expected message '%s', got '%s'", test.expectedMsg, resp.Message)
			} else if test.expectedErrMsg != "" && resp.Error != test.expectedErrMsg {
				t.Errorf("Expected error message '%s', got '%s'", test.expectedErrMsg, resp.Error)
			}
		})
	}
}
