package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type payload struct {
	ID        string `json:"id"`
	HeartRate int    `json:"heartRate"`
	Unit      string `json:"unit"`
	TimeStamp int    `json:"timeStamp"`
}

func main() {
	ticker := time.NewTicker(time.Millisecond * 5000)
	for range ticker.C {
		call()

	}
}

func call() {
	url := "http://129.146.106.151:8000/data"
	t := int(time.Now().Unix())
	m := payload{
		"001",
		80,
		"bpm",
		t,
	}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	fmt.Println("payload: ", m)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response: ", string(body))
}
