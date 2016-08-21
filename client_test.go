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
