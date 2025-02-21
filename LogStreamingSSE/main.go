package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	deploymentId = "123123"
)

func main() {
	fmt.Println("Starting to Server")
	StartServer()

	// client := &http.Client{}

	// startNewDeployment(client)

	// fmt.Println("Waiting for 5 sec")
	// time.Sleep(5 * time.Second)

	// fmt.Println("start event streaming")
	// startStreamDeploymentLogs(client)
}

func startNewDeployment(client *http.Client) {
	reqstruct := &StartDeploymentRequest{
		deploymentId: deploymentId,
	}

	reqByte, err := json.Marshal(reqstruct)
	if err != nil {
		panic(err)
	}
	b := bytes.NewBuffer(reqByte)

	req, err := http.NewRequest("GET", "http://localhost:8080/startDeployment", b)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting deployment")
	if _, err = client.Do(req); err != nil {
		panic(err)
	}
}

func fetchDeploymentLogs(client *http.Client) {
	fetchDepReq := &StartLogStreamingRequest{
		deploymentId: deploymentId,
	}

	reqByte, err := json.Marshal(fetchDepReq)
	if err != nil {
		panic(err)
	}
	b := bytes.NewBuffer(reqByte)

	fmt.Println("Starting to fetch deployment logs")
	req, err := http.NewRequest("GET", "http://localhost:8080/startLogStreaming", b)
	if err != nil {
		panic(err)
	}

	if _, err = client.Do(req); err != nil {
		panic(err)
	}
}

func startStreamDeploymentLogs(client *http.Client) {
	req, err := http.NewRequest("GET", "http://localhost:8080/streamDeploymentLogs", nil)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
		fmt.Println(res.Body)
	}
}
