package main

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestHelloWorldApi(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	result := `{"message":"You are visitor 0"}`
	// First call
	if assert.NoError(t, helloWorldEndPoint(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, result, strings.TrimSpace(rec.Body.String()))
	}

	rec = httptest.NewRecorder()
	context = e.NewContext(req, rec)
	result = `{"message":"You are visitor 1"}`

	// Second call
	if assert.NoError(t, helloWorldEndPoint(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, result, strings.TrimSpace(rec.Body.String()))
	}
}

