package decorator

type Sms struct {
	auth Authentication
}

func (s *Sms) login() bool {
	isValidSms := true
	return s.auth.login() && isValidSms
}
