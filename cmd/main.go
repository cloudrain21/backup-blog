package main

import (
	cm "github.com/cloudrain21/backup-blog/module/commander"
	"github.com/cloudrain21/backup-blog/module/logger"
	"log"
	"os"
)

func main() {
	logger.Init("./backup-blog.log")

	url := os.Args[1]
	user := os.Args[2]
	pass := os.Args[3]

	log.Printf("url : %s, user : %s/%s\n", url, user, pass)

	var commander cm.Commander
	commander = cm.NewSSH(url, user, pass)

	if err := commander.Connect(); err != nil {
		log.Fatal(err)
	}
	defer commander.Close()

	if err := commander.RunCommand("ls -al"); err != nil {
		log.Fatal(err)
	}

}
