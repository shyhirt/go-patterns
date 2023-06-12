package command

type Router struct {
	on bool
}

func (r Router) onOff() {
	r.on = !r.on
}
