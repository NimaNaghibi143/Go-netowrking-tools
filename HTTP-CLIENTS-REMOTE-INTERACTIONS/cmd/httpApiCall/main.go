package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Step #1
// Basic HTTP methods using the default http package functions.
// These are convenience wrappers around the default http.Client.

// func main() {
// 	r1, err := http.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer r1.Body.Close()
// 	fmt.Println("GET status:", r1.Status)

// 	r2, err := http.Head("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer r2.Body.Close()
// 	fmt.Println("HEAD status:", r2.Status)

// 	form := url.Values{}
// 	form.Add("foo", "bar")
// 	r3, err := http.Post(
// 		"https://www.google.com/robots.txt",
// 		"application/x-www-form-urlencoded",
// 		strings.NewReader(form.Encode()),
// 	)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer r3.Body.Close()
// 	fmt.Println("POST status:", r3.Status)

// 	// Alternative shorthand for form POST:
// 	// r3, err := http.PostForm("https://www.google.com/robots.txt", form)
// }

// Step #2
// Using a custom http.Client.
// The default client has no timeout — in a real tool this will hang forever
// on unresponsive hosts. A custom client gives you full control over
// timeouts, redirect policy, and transport-level settings (TLS, proxies).

// func main() {
// 	client := &http.Client{
// 		Timeout: 10 * time.Second,
// 	}

// 	resp, err := client.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Printf("Status : %s\n", resp.Status)
// 	fmt.Printf("Body   :\n%s\n", string(body))
// }

// Step #3
// Building a request manually with http.NewRequest.
// This gives full control over headers — useful for spoofing User-Agent,
// adding auth tokens, or crafting any custom HTTP interaction.

// func main() {
// 	req, err := http.NewRequest("GET", "http://www.google.com/robots.txt", nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; BlackHatGoBot/1.0)")

// 	client := &http.Client{Timeout: 10 * time.Second}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Printf("Status : %s\n", resp.Status)
// 	fmt.Printf("Body   :\n%s\n", string(body))
// }

// Step #4 (active)
// Combining all three methods with a shared custom client.
// Demonstrates GET, HEAD, and POST in a single run with proper
// error handling, response body draining, and timeout control.

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	// --- GET ---
	resp, err := client.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln("GET failed:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("reading GET body:", err)
	}
	fmt.Printf("[GET]  %s — %d bytes\n", resp.Status, len(body))

	// --- HEAD (no body, inspect headers only) ---
	resp2, err := client.Head("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln("HEAD failed:", err)
	}
	defer resp2.Body.Close()
	fmt.Printf("[HEAD] %s — Content-Type: %s\n", resp2.Status, resp2.Header.Get("Content-Type"))

	// --- POST with form data ---
	form := url.Values{}
	form.Add("foo", "bar")
	resp3, err := client.Post(
		"https://httpbin.org/post",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Fatalln("POST failed:", err)
	}
	defer resp3.Body.Close()
	body3, err := io.ReadAll(resp3.Body)
	if err != nil {
		log.Fatalln("reading POST body:", err)
	}
	fmt.Printf("[POST] %s — %d bytes\n", resp3.Status, len(body3))

	// --- Custom request with spoofed User-Agent ---
	req, err := http.NewRequest("GET", "http://www.google.com/robots.txt", nil)
	if err != nil {
		log.Fatalln("building request:", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; BlackHatGoBot/1.0)")

	resp4, err := client.Do(req)
	if err != nil {
		log.Fatalln("custom GET failed:", err)
	}
	defer resp4.Body.Close()
	fmt.Printf("[CUSTOM UA] %s\n", resp4.Status)

	_ = time.Now() // suppress unused import if steps above are commented out
}
