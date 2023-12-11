package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	fmt.Println("Test")
	tesseract := gosseract.NewClient()
	tesseract.SetImage("images/repo_screenshot.png")
	defer tesseract.Close()
}
