package command

type Sender struct {
	command Command
}

func (s *Sender) sendOnOff() {
	s.command.execute()
}
