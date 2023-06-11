package facade

type Account struct {
	login    string
	password string
}

func NewAccount(login, password string) *Account {
	return &Account{
		login:    login,
		password: password,
	}
}
