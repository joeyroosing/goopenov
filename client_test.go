package main

import (
	"fmt"
	"testing"
)

func TestGetDepartures(*testing.T) {
	fmt.Println("Starting!")
	a := NewAPIClient()
	resp, _ := a.GetDepartures("5839")
	fmt.Printf("Stop area contains: %+v", resp)
}

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
