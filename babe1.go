package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Define command-line flags
	phoneNumber := flag.String("number", "", "Phone number (e.g., 01922450716)")
	flag.Parse()

	// Check if the phone number flag is provided
	if *phoneNumber == "" {
		fmt.Println("Please provide a phone number using the -number flag.")
		os.Exit(1)
	}

	url := "https://www.beauti4me.com/api/v1/auth/resend-code"
	payload := []byte(fmt.Sprintf(`{"email":"","phone":"+88%s","code":"","invalidPhone":false,"showInvalidPhone":false}`, *phoneNumber))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", fmt.Sprint(len(payload)))
	req.Header.Set("Origin", "https://www.beauti4me.com")
	req.Header.Set("Referer", "https://www.beauti4me.com/user/verify-account")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Te", "trailers")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	// You can process the response body here if needed
}
