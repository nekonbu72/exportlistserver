package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	testServer := httptest.NewServer(newHandler())
	defer testServer.Close()

	r, err := http.Get(testServer.URL + "/ping")
	if err != nil {
		t.Errorf("http.Get(). %v", err)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(). %v", err)
		return
	}

	dataStr := string(data)
	if dataStr == "null\n" {
		// criteria や exportlist の列設定などのエラー
		t.Error("No data")
		return
	}
	if dataStr[2:7] == "error" {
		// メールのパスワードなど conninfo のエラー
		t.Error(dataStr)
		return
	}

	log.Printf("Response: \n%v\n", dataStr)
}
