package main

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	// Run Chrome browser
	pwd := "/Users/romelcampbell/GOCODE/LearingInTheWild/SeleniumWebDriverTest/BlogExample01"
	service, err := selenium.NewChromeDriverService(pwd+"/chromedriver", 4444)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		// "--headless",  // comment out this line to see the browser
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}

	driver.Get("https://www.google.com")
}
