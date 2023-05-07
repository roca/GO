package rgb_go_selenium

import (
	"io/ioutil"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"github.com/tebeka/selenium"
)

// Avoid Ginkgo and Gomega dot imports by assigning needed functions to variables.
var (
	Fail        = ginkgo.Fail
	RunSpecs    = ginkgo.RunSpecs
	Describe    = ginkgo.Describe
	BeforeEach  = ginkgo.BeforeEach
	AfterEach   = ginkgo.AfterEach
	It          = ginkgo.It
	CurrentTest = ginkgo.CurrentGinkgoTestDescription

	RegisterFailHandler = gomega.RegisterFailHandler
	Expect              = gomega.Expect
	HaveOccurred        = gomega.HaveOccurred
	BeZero              = gomega.BeZero
)

// URL returns full path for passed environment value.
func URL(env Env) string { return "http://" + urlMap[env] }

// TakeScreenshot saves screenshot of passed WebDriver into file with passed test name.
func TakeScreenshot(wd selenium.WebDriver, testName string) {
	bytes, err := wd.Screenshot()
	if err != nil {
		log.Panic().Err(err).Msg("Can't take a screenshot.")
	}
	ioutil.WriteFile(testName+".jpg", bytes, 0644)
}

// ErrCheck checks if error occurred.
func ErrCheck(err error) {
	Expect(err).ToNot(HaveOccurred())
}

// MustFindElement returns element or fails if element is not found.
func MustFindElement(wd selenium.WebDriver, by, value string) selenium.WebElement {
	element, err := wd.FindElement(by, value)
	ErrCheck(err)
	return element
}

// MustNotFindElement returns fails if element is found.
func MustNotFindElement(wd selenium.WebDriver, by, value string) {
	wd.SetImplicitWaitTimeout(time.Second)
	defer wd.SetImplicitWaitTimeout(DefTimeout)
	element, err := wd.FindElement(by, value)
	Expect(element).To(BeZero())
	Expect(err).To(HaveOccurred())
}

// MustWaitWithTimeout wait for passed selenium.Condition a given amount of time and checks for returned error value.
func MustWaitWithTimeout(wd selenium.WebDriver, condition selenium.Condition, timeout time.Duration) {
	ErrCheck(wd.WaitWithTimeout(condition, timeout))
}
