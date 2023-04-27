// Package api defines REST API /api.
package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Setup REST API
	router := gin.Default()
	apiRouter := router.Group("/api")
	apiRouter.GET("/healthchecker", Healthcheck)

	// Send request
	req, err := http.NewRequest("GET", "/api/healthchecker", nil)
	assert.Nil(t, err)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Assert result
	response, err := io.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, recorder.Code)
	assert.Empty(t, "", response)
}
