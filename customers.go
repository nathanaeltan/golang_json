package main

import (
	"encoding/json"
	"io/ioutil"
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

func readJsonAndReturnCustomers(filename string) (Customers, error) {
	plan, _ := ioutil.ReadFile(filename)

	var data Customers
	err := json.Unmarshal([]byte(plan), &data)

	if err != nil {
		return Customers{}, err
	}

	return data, nil
}
