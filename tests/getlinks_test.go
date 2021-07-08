package tests

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"linkscape/controllers"
)

func TestGetLinks(T *testing.T){
	req, err := http.NewRequest("GET", "/api/links", nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetAllLink)
	handler.ServeHTTP(rr, req)
}
