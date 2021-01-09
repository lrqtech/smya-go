package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// http请求
func Request(method string, url string, body []byte, limit int64) []byte {
	client := &http.Client{}
	switch method {
	case "get", "GET":
		method = "GET"
	case "post", "POST":
		method = "POST"
	}
	r, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(io.LimitReader(resp.Body, limit))
	if err != nil {
		fmt.Println(err)
	}
	return data
}
