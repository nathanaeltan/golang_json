package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
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

func makeRequest(c int) error {
	url := "https://notification.decathlon.sg/api/order/getOrderByCustomer.php"
	postValue := map[string]int{
		"customerId": c,
	}
	jsonValue, _ := json.Marshal(postValue)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	defer resp.Body.Close()
	bytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return err
	}
	var orders Orders
	err = json.Unmarshal(bytes, &orders)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	orders.print(c)
	return nil
}

func (o Orders) print(c int) {
	resultCustomerIdStr := fmt.Sprintln("---------------- Result of Customer ID:", c, "-----------------")
	fileData := []string{}

	fmt.Println(resultCustomerIdStr)
	fileData = append(fileData, resultCustomerIdStr)
	if len(o.Orders) > 0 {
		noOfOrdersStr := fmt.Sprintf("There are %v orders\n", o.NumberOfOrders)
		fmt.Println(noOfOrdersStr)
		fileData = append(fileData, noOfOrdersStr)
		for _, ord := range o.Orders {
			startOrderStr := fmt.Sprintln("----------------    Start of Order          -----------------")
			fmt.Println(startOrderStr)
			fileData = append(fileData, startOrderStr)

			// fmt.Printf("%+v\n\n", ord)
			v := reflect.ValueOf(ord)
			typeofV := v.Type()
			for i := 0; i < v.NumField(); i++ {
				orderFieldValueStr := fmt.Sprintln(typeofV.Field(i).Name, ":", v.Field(i).Interface())
				fmt.Println(orderFieldValueStr)
				fileData = append(fileData, orderFieldValueStr)

			}
			endOfOrderStr := fmt.Sprintln("----------------    End of Order            -----------------")
			fmt.Println(endOfOrderStr)
			fileData = append(fileData, endOfOrderStr)

		}
	} else {
		noOrdersStr := fmt.Sprintln("There are no Orders")
		fmt.Println(noOrdersStr)
		fileData = append(fileData, noOrdersStr)
	}
	endResultStr := fmt.Sprintln("----------------    End Result           -----------------")
	fmt.Println(endResultStr)
	fileData = append(fileData, endResultStr)
	err := writeToFile(strings.Join(fileData, ""))
	if err != nil {
		fmt.Println("Error in Writing to file: ", err)
	}
}

func writeToFile(text string) error {
	f, err := os.OpenFile("data", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}
