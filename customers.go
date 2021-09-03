package main

import (
	_ "embed"
	"encoding/json"
)

type Customers struct {
	Customers []struct {
		CustomerId int     `json:"customerId"`
		Country    string  `json:"country"`
		Language   string  `json:"language"`
		Age        int     `json:"age"`
		Weight     float32 `json:"weight"`
	} `json:"customers"`
}

//go:embed raw.json
var jsonFile string

func readJsonAndReturnCustomers() (Customers, error) {

	var data Customers

	err := json.Unmarshal([]byte(jsonFile), &data)

	if err != nil {
		return data, err
	}

	return data, nil
}
