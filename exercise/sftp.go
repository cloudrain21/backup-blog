package main

import (
	"bufio"
	"github.com/pkg/sftp"
	"github.com/sfreiberg/simplessh"
	"os"
	"syscall"
)

func GetFileSize(url string, user string, pass string, fileName string) (int64, error) {
	var err error
	var sshftp *sftp.Client

	conn, err := SshConnect(url, user, pass)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	if sshftp, err = sftp.NewClient(conn.SSHClient); err != nil {
		return 0, err
	}
	defer sshftp.Close()

	fi, err := sshftp.Lstat(fileName)
	if err != nil {
		return 0, err
	}

	return fi.Size(), nil
}

func DownloadFile(url string, user string, pass string, remoteFile string, saveFile string) (int64, error) {
	var sshftp *sftp.Client
	var conn *simplessh.Client
	var err error
	var fileInfo os.FileInfo
	var srcFile *sftp.File
	var trgFile *os.File

	// ssh connection
	if conn, err = SshConnect(url, user, pass); err != nil {
		return 0, err
	}
	defer conn.Close()

	// sftp client on ssh connection
	if sshftp, err = sftp.NewClient(conn.SSHClient); err != nil {
		return 0, err
	}
	defer sshftp.Close()

	// get file size to download
	if fileInfo, err = sshftp.Lstat(remoteFile); err != nil {
		return 0, err
	}
	fileSize := fileInfo.Size()

	// open file to download
	if srcFile, err = sshftp.OpenFile(remoteFile, syscall.O_RDONLY); err != nil {
		return 0, err
	}
	defer srcFile.Close()

	// read file into buffer
	buf := make([]byte, fileSize)
	bufr := bufio.NewReader(srcFile)

	_, err = bufr.Read(buf)
	if err != nil {
		return 0, err
	}

	// save data in buffer into local file
	if trgFile, err = os.OpenFile(saveFile, syscall.O_RDWR|syscall.O_CREAT, 0666); err != nil {
		return 0, err
	}
	defer trgFile.Close()

	bufw := bufio.NewWriter(trgFile)
	if _, err = bufw.Write(buf); err != nil {
		return 0, err
	}

	if fileInfo, err = os.Stat(saveFile); err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}
