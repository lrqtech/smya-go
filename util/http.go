package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// http请求
func Request(method string, url string, body []byte, header string) []byte {
	client := &http.Client{}
	switch method {
	case "get", "GET":
		method = "GET"
	case "post", "POST":
		method = "POST"
	case "put", "PUT":
		method = "PUT"
	case "delete", "DELETE":
		method = "DELETE"
	}
	r, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	b := []byte(header)
	h := make(map[string]string)
	err = json.Unmarshal(b, &h)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range h {
		r.Header.Add(k, v)
	}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
