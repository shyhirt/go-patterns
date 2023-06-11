package decorator

type Authentication interface {
	login() bool
}
