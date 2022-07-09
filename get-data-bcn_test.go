package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetDataBcn(t *testing.T) {
	mock := ``
	r := SetUpRouter()
	r.GET("/get-data-bcn", getDataBCN)
	req, _ := http.NewRequest("GET", "/get-data-bcn", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	response, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mock, string(response))
	assert.Equal(t, http.StatusOK, w.Code)

}
