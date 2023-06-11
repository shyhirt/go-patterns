package adapter

import "testing"

func TestDviAdapter_InsertIntoHDMI(t1 *testing.T) {
	user := &User{}
	tvHDMI := &TvHDMI{}
	user.PutHDMIonTV(tvHDMI)
	tvDVI := &TvDVI{}
	adapter := &DviAdapter{
		tvDVI,
	}
	user.PutHDMIonTV(adapter)
}
