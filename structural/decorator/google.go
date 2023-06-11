package decorator

type Google struct {
}

func (g *Google) login() bool {
	//authentication logic and get token
	validGoogleLoginPass := true
	return validGoogleLoginPass
}
