package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	deliveryEndpoint = "http://localhost:8081"
	foodEndpoint     = "http://localhost:8082"
	reserve          = "reserve"
	assign           = "assign"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	for i := range 10 {
		if err := placeOrder(client); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Successfully placed order", i)
		}

	}
}

func placeOrder(client *http.Client) error {

	_, err1 := client.Get(deliveryEndpoint + "/" + reserve)
	if err1 != nil {
		return errors.New("Could not reserve delivery agent")
	}

	_, err2 := client.Get(foodEndpoint + "/" + reserve)
	if err2 != nil {
		return errors.New("Could not reserve food")
	}

	_, err3 := client.Get(deliveryEndpoint + "/" + assign)
	if err3 != nil {
		return errors.New("Could not assign delivery agent")
	}

	_, err4 := client.Get(foodEndpoint + "/" + assign)
	if err4 != nil {
		return errors.New("Could not reserve food")
	}

	return nil
}
