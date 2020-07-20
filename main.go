package main

import (
	"fmt"
	"log"

	"github.com/sclevine/agouti"
)

func main() {

	url := "http://localhost:4444/wd/hub"
	options := []agouti.Option{agouti.Browser("chrome")}
	page, err := agouti.NewPage(url, options...)
	if err != nil {
		log.Fatalf("new page failed. %v\n", err)
	}

	err = page.Navigate("https://example.com")
	if err != nil {
		log.Fatalf("navigate example.com failed. %v\n", err)
	}

	title, err := page.Title()
	if err != nil {
		log.Fatalf("get title failed. %v\n", err)
	}
	fmt.Println(title)
}
