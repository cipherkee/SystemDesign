package main

/*
	/startDeployment - starts a go routine, and writes to a log file every 2s some lines of code
	/getDeploymentLogs - starts a SSE connection, and reads lines from log file and send in conn continuously
*/

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func StartServer() {
	http.HandleFunc("/startDeployment", StartDedployment)

	http.HandleFunc("/getLogs", getLogs)

	http.HandleFunc("/streamDeploymentLogs", streamDeploymentLogs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func StartDedployment(w http.ResponseWriter, r *http.Request) {

	req := StartDeploymentRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
	}

	fo, err := os.Create(fmt.Sprintf("logFile/%v_output.txt", req.deploymentId))
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			lines := fmt.Sprintf("[%v]: Writing new Line\n", time.Now())
			byteline := []byte(lines)
			if _, err := fo.Write(byteline); err != nil {
				panic(err)
			}
			time.Sleep(time.Second)
		}
	}()
}

func getLogs(w http.ResponseWriter, r *http.Request) {
	req := StartLogStreamingRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
	}

	fileName := fmt.Sprintf("logFile/%v_output.txt", req.deploymentId)

	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(fi)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	fmt.Println(fileLines)

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
}

func streamDeploymentLogs(w http.ResponseWriter, r *http.Request) {

	rc := http.NewResponseController(w)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	t := time.NewTicker(time.Second)
	clientGone := r.Context().Done()

	defer t.Stop()
	for {
		select {
		case <-clientGone:
			fmt.Println("client disconnected")
		case <-t.C:
			fmt.Fprintf(w, "data: Time is %v\n\n", time.Now())
			rc.Flush()
		}
	}
}
