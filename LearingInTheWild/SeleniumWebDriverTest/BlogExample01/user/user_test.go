package user

import (
	"fmt"
	"testing"
	"time"

	. "github.com/matijakrajnik/rgb_go_selenium"

	"github.com/BurntSushi/xgbutil"
	"github.com/rs/zerolog/log"
	"github.com/tebeka/selenium"
)

// User test suite.
func TestUser(t *testing.T) {
	ParseArgs()
	RegisterFailHandler(Fail)
	RunSpecs(t, "User")
}

var _ = Describe("User", func() {
	var (
		service *selenium.Service
		display *xgbutil.XUtil
		wd      selenium.WebDriver

		username = fmt.Sprintf("batman_%v", time.Now().Unix())
		password = "secret123"
	)

	BeforeEach(func() {
		var err error
		service, err = StartSelenium()
		Expect(service).ToNot(BeZero())
		ErrCheck(err)
		display, err = ConnectToDisplay()
		Expect(display).ToNot(BeZero())
		ErrCheck(err)
		wd, err = selenium.NewRemote(GetCaps(), fmt.Sprintf("http://localhost:%d/wd/hub", GetConf().Port))
		ErrCheck(err)
		Expect(wd).ToNot(BeZero())
		ErrCheck(wd.SetImplicitWaitTimeout(DefTimeout))
		ErrCheck(wd.Get(URL(GetConf().Env)))
	})

	AfterEach(func() {
		TakeScreenshot(wd, CurrentTest().TestText)
		err := wd.Quit()
		ErrCheck(err)
		display.Conn().Close()
		if err := service.Stop(); err != nil {
			log.Error().Err(err).Msg("Error while stoping Selenium server.")
		}
	})

	It("can create new account", func() {
		loginLink := MustFindElement(wd, selenium.ByLinkText, "LOGIN")
		ErrCheck(loginLink.Click())
		newAccountLink := MustFindElement(wd, selenium.ByCSSSelector, ".btn-link")
		ErrCheck(newAccountLink.Click())
		usernameInput := MustFindElement(wd, selenium.ByID, "username")
		ErrCheck(usernameInput.SendKeys(username))
		passwordInput := MustFindElement(wd, selenium.ByID, "password")
		ErrCheck(passwordInput.SendKeys(password))
		submitButton := MustFindElement(wd, selenium.ByCSSSelector, ".btn-success")
		ErrCheck(submitButton.Click())
		MustWaitWithTimeout(wd, func(wd selenium.WebDriver) (bool, error) {
			header := MustFindElement(wd, selenium.ByTagName, "h1")
			text, err := header.Text()
			return text == "Welcome to React Gin Blog!", err
		}, 5*time.Second)
		logoutLink := MustFindElement(wd, selenium.ByCSSSelector, ".btn-dark")
		Expect(logoutLink).ToNot(BeZero())
	})
})
