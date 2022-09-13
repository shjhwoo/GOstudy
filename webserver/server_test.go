package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHnadler(t *testing.T){
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)


	assert.Equal(http.StatusOK,res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("hello world", string(data))
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK,res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("hello bar", string(data))
}