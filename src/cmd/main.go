package main

import (
	"fmt"
	"log"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	fmt.Println("Hello from a Go OCR client!")

	tessClient := gosseract.NewClient()
	defer tessClient.Close()
	tessClient.SetImage("./images/repo_screenshot.png")
	text, err := tessClient.Text()
	if err != nil {
		log.Println(err)
	}
	log.Printf("Recognized text:\n%v\n", text)

	tessClient.SetImage("./images/repo_screenshot_2.jpg")
	text, err = tessClient.Text()
	if err != nil {
		log.Println(err)
	}
	log.Printf("Recognized text:\n%v\n", text)
}
