package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Define command-line flag
	phoneNumber := flag.String("number", "", "Phone number (e.g., 01922450716)")
	flag.Parse()

	// Check if the phone number flag is provided
	if *phoneNumber == "" {
		fmt.Println("Please provide a phone number using the -number flag.")
		os.Exit(1)
	}

	// Remove the first digit from the phone number
	trimmedNumber := (*phoneNumber)[1:]

	url1 := "https://mygirlco.com/api/web"

	// Set the request payload
	payload := url.Values{
		"phone":        {trimmedNumber},
		"country_code": {"880"},
		"action":       {"send_otp"},
	}

	// Create the HTTP client
	client := &http.Client{}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url1, strings.NewReader(payload.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the request headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Referer", "https://mygirlco.com/login")
	req.Header.Set("Origin", "https://mygirlco.com")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Te", "trailers")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

}
