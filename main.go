package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://v0.ovapi.nl/stopareacode/5839/", nil)
	req.Header.Add("Content-Encoding", "gzip")
	resp, err := client.Do(req)
	if err != nil {
		panic("Could not retrieve data")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Could not retrieve data")
	}

	var stopareas map[string]interface{}
	err2 := json.Unmarshal(body, &stopareas)
	if err2 != nil {
		panic("Could not parse json data")
	}
	res, _ := json.MarshalIndent(stopareas, "", "\t")
	fmt.Println(string(res))
}
