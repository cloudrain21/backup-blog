package main

import (
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("log/backup_homepage.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
}

