package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBook(t *testing.T) {
	srv := NewServer()
	testSrv := httptest.NewServer(http.HandlerFunc(srv.GetBook))
	num := 1000

	for i := 0; i < num; i++ {
		id := i%100 + 1
		url := fmt.Sprintf("%s/?id=%d", testSrv.URL, id)
		resp, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}
		book := &Book{}
		if err := json.NewDecoder(resp.Body).Decode(book); err != nil {
			t.Error(err)
		}

		fmt.Printf("%+v\n", book)
	}
	fmt.Printf("The database was called %d times: ", srv.count)
}
