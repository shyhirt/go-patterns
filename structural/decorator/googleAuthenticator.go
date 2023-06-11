package decorator

type GoogleAuthenticator struct {
	auth Authentication
}

func (g *GoogleAuthenticator) login() bool {
	isValidSms := true
	return g.auth.login() && isValidSms
}
