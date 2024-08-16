package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Ricid",
		MiddleName: "Kumbara",
		LastName:   "Kagenou",
		Age:        25,
		Married:    false,
	}

	byte, _ := json.Marshal(customer)
	fmt.Println(string(byte))
}

func TestDecode(t *testing.T) {
	jsonRequest := `{"FirstName":"Ricid","MiddleName":"Kumbara","LastName":"Kagenou","Age":25,"Married":false}`
	jsonByte := []byte(jsonRequest)

	customer := &Customer{}
	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(*customer)
	fmt.Println(&customer)
	fmt.Println(customer.FirstName)
}
