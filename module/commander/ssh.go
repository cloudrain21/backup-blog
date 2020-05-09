package commander

import (
	"errors"
	"github.com/sfreiberg/simplessh"
	"log"
)

type SSH struct {
	url  string
	user string
	pass string
	conn *simplessh.Client
}

var (
	ErrSSHConnect = errors.New("ssh connection error")
)

func (s *SSH) Connect() error {
	log.Println("connect...")

	client, err := simplessh.ConnectWithPassword(s.url, s.user, s.pass)
	if err != nil {
		return ErrSSHConnect
	}

	s.conn = client
	return nil
}

func (s SSH) RunCommand(command string) error {
	log.Printf("run command : %s\n", command)
	return nil
}

func NewSSH(url string, user string, pass string) *SSH {
	return &SSH{url, user, pass, nil}
}
