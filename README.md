# Screenshot Capture API Go Client

Webseite-Herunterladen.de Screenshot Capture is a very simple but powerful screenshot API that anyone can easily use to create pixel-perfect website screenshots. It always uses a recent version of Chrome to ensure that all modern web features are fully supported and rendering is exactly as your customers would expect.

```go
package main

import (
	"fmt"
	"github.com/Webseite-Herunterladen-de/screenshot-capture-api-go/client"
	"github.com/Webseite-Herunterladen-de/screenshot-capture-api-go/client/screenshot"
	"os"
)

func main() {

	//define params
	params := screenshot.NewCaptureScreenshotUnauthenticatedParams()

	//set url and token
	params.URL = "https://webseite-herunterladen.de"
	params.Token = "YOUR_TOKEN"

	//create file
	f, err := os.Create("webseite_herunterladen_screenshot.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	//send request
	response, err := client.Default.Screenshot.CaptureScreenshotUnauthenticated(params, f)
	if err != nil {
		fmt.Println(err)
		return
	}

	//check for response errors and capture points
	if response.XREMAININGQUOTA + response.XREMAININGQUOTAPREPAID == 0 {
		fmt.Println("Oh no! No quota left.")
	}
}

```