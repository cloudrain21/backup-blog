package main

import (
	"errors"
	"log"
	"time"
)

type Downloader struct {
	url string
	user string
	pass string
	remotePath string
	localPath string
	tarFileName string
}

var (
	InvalidDownloadInfo = errors.New("invalid download information")
	InvalidFileSize = errors.New("download fail. different file size")
)

func NewDownloader(url string, user string, pass string) *Downloader {
	return &Downloader{
		url,
		user,
		pass,
		"",
		"",
		"",
	}
}

func (d *Downloader)SetremotePath(path string) {
	d.remotePath = path
}

func (d *Downloader)SetLocalPath(path string) {
	d.localPath = path
}

func (d *Downloader)Download() error {
	if err := d.makeTgz(); err != nil {
		return err
	}

	if err := d.downloadTgz(); err != nil {
		return err
	}
	return nil
}

func (d *Downloader)downloadTgz() error {
	var srcFileSize, trgFileSize int64 = 0, 0
	var err error

	if srcFileSize, err = GetFileSize(d.url, d.user, d.pass, d.remotePath + "/" + d.tarFileName); err != nil {
		return err
	}

	remoteFile := d.remotePath + "/" + d.tarFileName
	saveFile := d.localPath + "/" + d.tarFileName

	if trgFileSize, err = DownloadFile(d.url, d.user, d.pass, remoteFile, saveFile); err != nil {
		return err
	}

	if srcFileSize == 0 || trgFileSize == 0 || srcFileSize != trgFileSize {
		return InvalidFileSize
	}
	return nil
}

func (d *Downloader)makeTgz() error {
	if err := d.validateDownloadInfo(); err != nil {
		return err
	}

	t := time.Now()
	dateString := t.Format("20060102150405")

	d.tarFileName = "wordpress_" + dateString + ".tgz"
	tarCommand := "tar cvzf " + d.tarFileName + " " + d.remotePath

	command := "cd " + d.remotePath + " ; rm -rf *.tgz ; " + tarCommand + " ; " + "exit"
	log.Printf("command : %s\n", command)

	if err := SshRunCommand(d.url, d.user, d.pass, command); err != nil {
		return err
	}
	return nil
}

func (d *Downloader)validateDownloadInfo() error {
	if d.url == "" || d.user == "" || d.pass == "" {
		return InvalidDownloadInfo
	}

	if d.remotePath == "" || d.localPath == "" {
		return InvalidDownloadInfo
	}

	return nil
}

