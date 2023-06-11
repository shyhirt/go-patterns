package adapter

import "log"

type DviAdapter struct {
	tv *TvDVI
}

func (t *DviAdapter) PutIntoHDMI() {
	log.Println("Adapter to HDMI -> DVI")
	//Adapter code in realisation
	t.tv.PutIntoDVI()
}
