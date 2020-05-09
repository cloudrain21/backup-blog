package main

import (
	"github.com/sfreiberg/simplessh"
)

type access struct {
	login    string
	password string
}


func SshRunCommand(url string, user string, pass string, command string) (error) {
	var client *simplessh.Client
	var err error

	if client, err = SshConnect(url, user, pass); err != nil {
		return err
	}
	defer client.Close()

	if _, err := client.Exec(command); err != nil {
		return err
	}
	return err
}

func SshConnect(url string, user string, pass string) (*simplessh.Client, error) {
	var client *simplessh.Client
	var err error

	client, err = simplessh.ConnectWithPassword(url, user, pass)
	if err != nil {
		return nil, err
	}

	return client, nil
}