package server

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/zfz7/tetris_go/backend/repository"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)


func TestHelloWorldApi(t *testing.T) {
	//Reset DB
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5433")
	os.Setenv("POSTGRES_PASSWORD", "tetris_test")
	os.Setenv("POSTGRES_USER", "tetris_test")
	os.Setenv("POSTGRES_DB", "tetris_test")
	repository.StartDB()
	var db, _ = repository.GetDB()
	db.Exec("TRUNCATE TABLE meta_data;")
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

