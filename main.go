package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
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

func main() {
	plan, _ := ioutil.ReadFile("raw.json")

	var data Customers
	err := json.Unmarshal([]byte(plan), &data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Customers)

	for _, customer := range data.Customers {
		makeRequest(customer.CustomerId)
	}

}

type Orders struct {
	NumberOfOrders int `json:"numberOfOrder"`
	Orders         []struct {
		OrderId      int    `json:"orderId"`
		CustomerId   int    `json:"customerId"`
		CustomerName string `json:"customerName"`
		TotalPaid    string `json:"totalPaid"`
		StoreId      int    `json:"storeId"`
	} `json:"orders"`
}

func makeRequest(c int) {
	url := "https://notification.decathlon.sg/api/order/getOrderByCustomer.php"
	postValue := map[string]int{
		"customerId": c,
	}
	jsonValue, _ := json.Marshal(postValue)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer((jsonValue)))
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	bytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		os.Exit(1)
	}
	var orders Orders
	err = json.Unmarshal(bytes, &orders)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("---------------- Result of Customer ID:", c, "-----------------")
	if len(orders.Orders) > 0 {
		fmt.Printf("There are %v orders\n", orders.NumberOfOrders)

		for _, ord := range orders.Orders {
			fmt.Println("----------------    Start of Order          -----------------")

			// fmt.Printf("%+v\n\n", ord)
			v := reflect.ValueOf(ord)
			typeofV := v.Type()
			for i := 0; i < v.NumField(); i++ {
				fmt.Println(typeofV.Field(i).Name, ":", v.Field(i).Interface())

			}

			fmt.Println("----------------    End of Order            -----------------")

		}
	} else {
		fmt.Println("There are no Orders")
	}
	fmt.Println("----------------    End Result           -----------------")

}
