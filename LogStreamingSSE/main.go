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
	go StartServer()

	client := &http.Client{}

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

	fmt.Println("Waiting for 5 sec")
	time.Sleep(10 * time.Second)

	fetchDepReq := &StartLogStreamingRequest{
		deploymentId: deploymentId,
	}

	reqByte, err = json.Marshal(fetchDepReq)
	if err != nil {
		panic(err)
	}
	b = bytes.NewBuffer(reqByte)

	fmt.Println("Starting to fetch deployment logs")
	req, err = http.NewRequest("GET", "http://localhost:8080/startLogStreaming", b)
	if err != nil {
		panic(err)
	}

	if _, err = client.Do(req); err != nil {
		panic(err)
	}
}
