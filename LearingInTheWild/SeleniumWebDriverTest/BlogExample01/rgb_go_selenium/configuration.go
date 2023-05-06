package rgb_go_selenium

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
)

// Env represents environment domain where tests will be run.
type Env string

// Browser represents browser type in which tests will be run.
type Browser string

const (
	// Env keys.
	DevEnv     Env = "dev"
	UATEnv     Env = "uat"
	PreprodEnv Env = "preprod"

	// Browser types.
	Chrome  Browser = "chrome"
	Firefox Browser = "firefox"

	// Paths to necessarry binaries. Chenge these to match to binary locations on your machine.
	seleniumPath     = "/usr/share/java/selenium-server.jar"
	geckoDriverPath  = "/usr/bin/geckodriver"
	chromeDriverPath = "/usr/bin/chromedriver"

	// Default timeout for WebDriver.
	DefTimeout = 5 * time.Second
)

var urlMap = map[Env]string{
	DevEnv:     "localhost:8080",
	UATEnv:     "uat.rgb.com",
	PreprodEnv: "preprod.rgb.com",
}

// Conf represents configuration data.
type Conf struct {
	Browser        Browser
	Env            Env
	Headless       bool
	DisplayAddress string
	Port           int
	Width          int
	Height         int
}

var (
	conf Conf
	caps selenium.Capabilities
)

// GetConf returns current set configuration.
func GetConf() Conf { return conf }

// SetCaps defines Selenium capabailities based on passed configuration.
func SetCaps(cnf Conf) {
	switch cnf.Browser {
	case Firefox:
		setFirefoxCaps(cnf)
	case Chrome:
		setChromeCaps(cnf)
	default:
		log.Panic().Str("Browser", string(cnf.Browser)).Msg("Invalid Browser type.")
	}
}

// GetCaps returns currently set Selenium capabilities.
func GetCaps() selenium.Capabilities { return caps }

func setFirefoxCaps(cnf Conf) {
	args := []string{
		fmt.Sprintf("--width=%d", cnf.Width),
		fmt.Sprintf("--height=%d", cnf.Height),
	}
	if cnf.Headless {
		args = append(args, "-headless")
	}
	firefoxCaps := firefox.Capabilities{
		Args: args,
	}
	caps = selenium.Capabilities{
		"browserName":           "firefox",
		firefox.CapabilitiesKey: firefoxCaps,
	}
}

func setChromeCaps(cnf Conf) {
	args := []string{
		fmt.Sprintf("--window-size=%d,%d", cnf.Width, cnf.Height),
		"--ignore-certificate-errors",
		"--disable-extensions",
		"--no-sandbox",
		"--disable-dev-shm-usage",
	}
	if cnf.Headless {
		args = append(args, "--headless", "--disable-gpu")
	}
	chromeCaps := map[string]interface{}{
		"excludeSwitches": [1]string{"enable-automation"},
		"args":            args,
	}
	caps = selenium.Capabilities{
		"browserName":   "chrome",
		"chromeOptions": chromeCaps,
	}
}
