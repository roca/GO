package rgb_go_selenium

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

// Define CLI arguments with name, default value and description.
var (
	browser        = flag.String("browser", "chrome", `Browser to run tests in. Possible values are "chrome" and "firefox"`)
	env            = flag.String("env", "dev", `Sets run environment. Possible values are "dev", "uat" and "preprod"`)
	headless       = flag.String("headless", "false", `Sets headless mode. Possible values are "false" and "true"`)
	displayAddress = flag.String("displayAddress", "", `X server address.`)
	port           = flag.Int("port", 4444, `Selenium server port. Must be a number between 1024-65535.`)
	width          = flag.Int("width", 1920, `Display width.`)
	height         = flag.Int("height", 1080, `Display height.`)
)

func usage() {
	fmt.Print(`This program runs RGB tests.
 
Usage:
 
go test [arguments]
 
Supported arguments:
 
`)
	flag.PrintDefaults()
}

// Parses passed arguments, sets conf and caps global variables.
func ParseArgs() {
	// Set function to be called if parsing fails.
	flag.Usage = usage

	// Parse CLI arguments.
	flag.Parse()

	// Print usage text and exit if:
	// - browser is neither "chrome" or "firefox",
	// - env is neither "dev", "uat" or "preprod",
	// - headless is neither "false" or "true",
	// - displayAddress is not valid IP address,
	// - port is not a number between 1024-65535
	isHeadless, err := strconv.ParseBool(*headless)
	if !(validBrowserArg() && validEnvArg() && err == nil && validDisplayArg() && (*port >= 1024 && *port <= 65535)) {
		usage()
		os.Exit(2)
	}

	// Set conf global variable.
	conf = Conf{
		Browser:        Browser(*browser),
		Env:            Env(*env),
		Headless:       isHeadless,
		DisplayAddress: *displayAddress,
		Port:           *port,
		Width:          *width,
		Height:         *height,
	}

	// Set caps global variable.
	SetCaps(conf)
}

func validBrowserArg() bool {
	return (*browser) == string(Chrome) || *browser == string(Firefox)
}

func validEnvArg() bool {
	return *env == string(DevEnv) || *env == string(UATEnv) || *env == string(PreprodEnv)
}

func validDisplayArg() bool {
	_, err := url.Parse(*displayAddress)
	return err == nil
}
