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
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
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
	orders.print(c)

}



func (o Orders) print(c int) {
	fmt.Println("---------------- Result of Customer ID:", c, "-----------------")
	if len(o.Orders) > 0 {
		fmt.Printf("There are %v orders\n", o.NumberOfOrders)

		for _, ord := range o.Orders {
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