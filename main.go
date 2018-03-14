package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type payload struct {
	ID        string `json:"id"`
	HeartRate int    `json:"heartRate"`
	Unit      string `json:"unit"`
	TimeStamp int    `json:"timeStamp"`
}

func main() {
	var count = 0
	ticker := time.NewTicker(time.Millisecond * 1000)
	for range ticker.C {
		call()
		count++
		if count >= 5 {
			os.Exit(0)
		}
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
