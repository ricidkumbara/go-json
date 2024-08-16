package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestStreamEncoder(t *testing.T) {
	w, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(w)

	customer := Customer{
		FirstName:  "Ricid",
		MiddleName: "Kumbara",
		LastName:   "Kagenou",
	}

	encoder.Encode(customer)
}
