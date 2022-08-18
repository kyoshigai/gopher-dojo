package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// TRY 【TRY】リクエストを送ってみよう
	v := url.Values{"p": []string{"Gopher"}}
	req, err := http.NewRequest("GET", "http://localhost:8082?"+v.Encode(), nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	byteStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(byteStr))
}
