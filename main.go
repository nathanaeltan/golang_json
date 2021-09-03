package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := readJsonAndReturnCustomers()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if len(data.Customers) > 0 {
		for _, customer := range data.Customers {
			err = makeRequest(customer.CustomerId)
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("No Customers found. Exiting Program ....")
		os.Exit(1)
	}

}
