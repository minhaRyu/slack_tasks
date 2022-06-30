package webdriverTest

import (
	"log"
	"time"

	"github.com/fedesog/webdriver"
)

func Test() {
	chromeDriver := webdriver.NewChromeDriver("./webdriverTest/chromedriver.exe")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}

	desired := webdriver.Capabilities{"Platform": "Windows"}
	required := webdriver.Capabilities{}
	session, _ := chromeDriver.NewSession(desired, required)
	err = session.Url("https://flex.team/auth/login")
	if err != nil {
		log.Println(err)
	}

	// ('[data-key="auth.login_flow.google_login_cta"]')
	// ('[data-identifier$="@wonderwall.kr"]')
	loginEl, _ := session.FindElement("css selector", "[data-key=\"auth.login_flow.google_login_cta\"]")
	loginEl.Click()
	log.Println(loginEl)
	time.Sleep(10 * time.Second)
	loginEl, _ = session.FindElement("css selector", "[data-identifier$=\"@wonderwall.kr\"]")
	loginEl.Click()
	log.Println(loginEl)
}
