package command

type Command struct {
	device Device
}

func (c *Command) execute() {
	c.device.onOff()
}
