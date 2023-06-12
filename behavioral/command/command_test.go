package command

import "testing"

func TestCommand(t *testing.T) {
	//On/off router
	router := Router{}
	onOffCommand := Command{router}
	onOffSender := &Sender{onOffCommand}
	onOffSender.command.execute()
}
