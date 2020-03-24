package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sugan2111/Intercom/models"
	"io"
	"net/http"
	"sort"
	"strconv"
)

func UrlToLines(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return LinesFromReader(resp.Body)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func GetCustomersList(custData []string, customer models.Customer) {
	for _, eachline := range custData {

		err := json.Unmarshal([]byte(eachline), &customer)
		if err != nil {
			fmt.Println("error:", err)
		}
		custLatitude, err := strconv.ParseFloat(customer.Latitude, 64)
		if err != nil {
			fmt.Println("error:", err)
		}

		custLongitude, err := strconv.ParseFloat(customer.Longitude, 64)
		if err != nil {
			fmt.Println("error:", err)
		}

		distance := CalculateDistance(53.339428, -6.257664, custLatitude, custLongitude)
		getInvitees(distance, customer)
	}
}

func getInvitees(distance float64, customer models.Customer) {
	if distance <= 100.0 {
		m := make(map[int]string)

		m[customer.UserId] = customer.Name // Add a new key-value pair

		// To store the keys in slice in sorted order
		var keys []int
		for k := range m {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		// To perform the opertion you want
		for _, k := range keys {
			fmt.Println(k, m[k])
		}
	}
}
