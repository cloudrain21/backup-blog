package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// mysql : mysql --host=db.cloudrain21.com --port=3306 -ucloudrain21 -p

func main() {
	if len(os.Args) < 2 {
		showUsage(os.Args)
	}

	url := os.Args[1]
	user := os.Args[2]
	pass := os.Args[3]

	// tar/gzip wordpress files/dirs and download it
	downloader := NewDownloader(url, user, pass)
	downloader.SetremotePath("/www_root")
	downloader.SetLocalPath("./back")

	if err := downloader.Download(); err != nil {
		log.Fatal(err)
	}

	// mysql data dump
	dataSourceName := "tcp:db.cloudrain21.com:3306*dbcloudrain21/" + user + "/" + pass
	if err := DumpDB(dataSourceName, "./back"); err != nil {
		log.Fatal(err)
	}

	// upload back files into google drive
	// https://gist.github.com/TheGU/e6d0ae13f2fa83f3bd8d
}

func showUsage(args []string) {
	fmt.Printf("Usage : %s homepage_url user passwd\n", path.Base(args[0]))
	os.Exit(0)
}
