package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// Define command-line flags
	url := flag.String("url", "https://api.meenaclick.com/api/front/sms/send/pin", "API URL")
	cellPhone := flag.String("phone", "", "Cell phone number")
	flag.Parse()

	if *cellPhone == "" {
		fmt.Println("Phone number is required")
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomType := strconv.Itoa(rand.Intn(100))

	payload := []byte(fmt.Sprintf(`{"cell_phone":"%s","type":"%s"}`, *cellPhone, randomType))

	req, err := http.NewRequest("POST", *url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Host", "api.meenaclick.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0%00")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", "42")
	req.Header.Set("Origin", "http://www.meenaclick.com")
	req.Header.Set("Referer", "http://www.meenaclick.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Te", "trailers")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response body:", string(body))
}
