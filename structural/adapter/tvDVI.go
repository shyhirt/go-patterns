package adapter

import (
	"log"
)

type TvDVI struct {
}

func (t *TvDVI) PutIntoDVI() {
	log.Println("HDMI on DVI supported TV")
}
