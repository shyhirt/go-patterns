package decorator

type Facebook struct {
}

func (f *Facebook) login() bool {
	//authentication logic and get token
	validFacebookLoginPass := true
	return validFacebookLoginPass
}
