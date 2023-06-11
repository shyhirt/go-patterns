package adapter

import "fmt"

type TvHDMI struct {
}

func (t *TvHDMI) PutIntoHDMI() {
	fmt.Println("HDMI supported TV")
}
