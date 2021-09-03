package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := readJsonAndReturnCustomers("raw.json")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	for _, customer := range data.Customers {
		makeRequest(customer.CustomerId)
	}

}

