package main

import (
	"testing"
)

func TestReadJsonAndReturnCustomers(t *testing.T) {
	customers, _ := readJsonAndReturnCustomers()
	if len(customers.Customers) != 3 {
		t.Error("Expected only 3 customers, instead received", len(customers.Customers))
	}

}
