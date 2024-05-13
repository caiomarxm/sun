package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "The name to greet.")
	flag.Parse()
	url := buildUrl()
	fmt.Printf("Hello, %s!\n", *name)
	fmt.Printf("Hello, %s!\n", url)
}

func buildUrl() string {
	base_url := os.Getenv("BASE_URL")
	api_key := os.Getenv("API_KEY")
	return base_url + "current.json?key=" + api_key
}
