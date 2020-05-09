package main

import (
	"database/sql"
	"errors"
	"github.com/JamesStewy/go-mysqldump"
	"github.com/ziutek/mymysql/godrv"
)

var (
	ErrorOpenDatabase = errors.New("failed to open database")
	ErrorRegisterDatabase = errors.New("failed to register database")
	ErrorDumpData = errors.New("failed to dump database")
)

func DumpDB(dataSourceName string, backupDir string) error {
	var dumper *mysqldump.Dumper
	var err error
	var db *sql.DB

	// Register the mymysql driver
	godrv.Register("SET NAMES utf8mb4")

	// Open connection to database
	if db, err = sql.Open("mymysql", dataSourceName); err != nil {
		return ErrorOpenDatabase
	}

	// Register database with mysqldump
	if dumper, err = mysqldump.Register(db, backupDir, "20060102150405"); err != nil {
		return ErrorRegisterDatabase
	}

	// Dump database to file
	if _, err = dumper.Dump(); err != nil {
		return ErrorDumpData
	}

	// Close dumper and connected database
	dumper.Close()
	return nil
}
