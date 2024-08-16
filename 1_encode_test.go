package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestEncodeSimple(t *testing.T) {
	logJson("Ricid")
	logJson(1)
	logJson(true)
	logJson([]string{"Ricid", "Kumbara", "Fulan"})
}
