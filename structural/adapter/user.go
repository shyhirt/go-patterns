package adapter

import "log"

type User struct {
}

func (u *User) PutHDMIonTV(tv TV) {
	log.Println("Put hdmi cable to different type TV")
	tv.PutIntoHDMI()
}
