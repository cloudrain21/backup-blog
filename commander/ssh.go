package commander

type SSH struct {
	url  string
	user string
	pass string
}

func (s *SSH) Connect() error {
	return nil
}

func (s SSH) RunCommand(command string) error {
	return nil
}

func NewSSH(url string, user string, pass string) *SSH {
	return &SSH{url, user, pass}
}
