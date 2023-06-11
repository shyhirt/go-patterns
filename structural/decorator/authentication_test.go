package decorator

import (
	"log"
	"testing"
)

func TestAuthentication(t *testing.T) {
	// 2FA
	loginFacebook := &Facebook{}
	loginFacebookWithSms := &Sms{
		auth: loginFacebook,
	}
	log.Println(loginFacebookWithSms.login())
	// 3FA
	loginGoogle := &Google{}
	loginGoogleWithSms := &Sms{
		auth: loginGoogle,
	}
	loginGoogleWithAuthenticator := &GoogleAuthenticator{
		loginGoogleWithSms,
	}
	log.Println(loginGoogleWithAuthenticator.login())

}
