package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type HandlerTestCase struct {
	Name             string
	Method, URL      string
	Body             string
	Header           http.Header
	ExpectedStatus   int
	ExpectedResponse string
}

func ExecuteHandlerTestCase(t *testing.T, h http.HandlerFunc, tc HandlerTestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		var body io.Reader
		if tc.Body != "" {
			body = strings.NewReader(tc.Body)
		}

		req, err := http.NewRequest(tc.Method, tc.URL, body)
		assert.NoError(t, err)

		if tc.Header != nil {
			req.Header = tc.Header
		}

		rr := httptest.NewRecorder()
		h.ServeHTTP(rr, req)

		assert.Equal(t, tc.ExpectedStatus, rr.Code)
		if tc.ExpectedResponse != "" {
			assert.JSONEq(t, tc.ExpectedResponse, rr.Body.String())
		}
	})
}
