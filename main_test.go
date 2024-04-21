package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The following function makes the text where the number of cafes is greater than the number of cafes in the cafeList.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, responseRecorder.Code, http.StatusOK, "Expected status code 200")

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, responseRecorder.Body, "Body is empty")
	lenBody := len(strings.Split(body, ","))
	assert.Equal(t, lenBody, totalCount, "Expected 4 cafe")
}

// The following function makes the test in which the customer has written the wrong city.
func TestMainHandlerInWhichWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=minsk", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest, "Expected status code 400")

}
