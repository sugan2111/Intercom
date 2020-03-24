package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sugan2111/Intercom/models"
	"log"
	"math"
	"os"
	"strconv"
)

func calculateDistance(latitudeFrom, longitudeFrom, latitudeTo, longitudeTo float64) float64 {
	long1 := longitudeFrom * (math.Pi / 180)
	long2 := longitudeTo * (math.Pi / 180)
	lat1 := latitudeFrom * (math.Pi / 180)
	lat2 := latitudeTo * (math.Pi / 180)

	//Haversine Formula
	dlong := long2 - long1
	dlati := lat2 - lat1

	val := math.Pow(math.Sin(dlati/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlong/2), 2)

	res := 2 * math.Asin(math.Sqrt(val))

	radius := 6371

	return res * float64(radius)

}

func readCustomerFile() []string {
	readFile, err := os.Open("customers.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	return fileTextLines

}

func main() {
	/*url := "https://s3.amazonaws.com/intercom-take-home-test/customers.txt"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}*/
	var customer models.Customer
	custData := readCustomerFile()
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

		distance := calculateDistance(53.339428, -6.257664, custLatitude, custLongitude)
		if distance <= 100 {
			m := make(map[int]string)

			m[customer.UserId] = customer.Name            // Add a new key-value pair
			// To store the keys in slice in sorted order
		/*	var keys []int
			for k := range m {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			fmt.Println(keys)


			for _, k := range keys {
				fmt.Println(k, m[k])
			}

		*/

			for key, value := range m { // Order not specified
				fmt.Printf("%d \t %s\n", key, value)
			}

		}
	}

}
