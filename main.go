package main

import (
	"fmt"
	"github.com/Webseite-Herunterladen-de/screenshot-capture-api-go/client/screenshot"
	"os"
)

func main() {
	//init Client
	client := screenshot.New()

	//define params
	params := screenshot.CaptureScreenshotUnauthenticatedParams{}

	//set url and token
	params.URL = "https://webseite-herunterladen.de"
	params.Token = "1337"

	//create file
	f, err := os.Create("webseite_herunterladen_screenshot.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	//send request
	response, err := client.CaptureScreenshotUnauthenticated(&params, f)
	if err != nil {
		fmt.Println(err)
		return
	}

	//check for response errors and capture points
	if response.XREMAININGQUOTAPREPAID == 0 {
		fmt.Println("Oh no! No prepaid quota left.")
	}
}
