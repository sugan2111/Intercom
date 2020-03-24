package main

import (
	"fmt"
	"github.com/sugan2111/Intercom/models"
	"github.com/sugan2111/Intercom/services"
)

func main() {
	url := "https://s3.amazonaws.com/intercom-take-home-test/customers.txt"
	custData, err := services.UrlToLines(url)
	if err != nil {
		fmt.Println("error:", err)
	}
	var customer models.Customer
	services.GetCustomersList(custData, customer)
}
