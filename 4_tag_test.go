package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Asus",
		ImageUrl: "http://localhost",
	}

	byte, _ := json.Marshal(product)
	fmt.Println(string(byte))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
