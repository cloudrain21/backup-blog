package main

import (
	"database/sql"
	"fmt"
	"github.com/JamesStewy/go-mysqldump"
	"github.com/ziutek/mymysql/godrv"
	"time"
)

func main() {
	// Register the mymysql driver
	//godrv.Register("SET NAMES utf8")
	godrv.Register("SET NAMES utf8mb4")

	// Open connection to database
	db, err := sql.Open("mymysql", "tcp:db.cloudrain21.com:3306*dbcloudrain21/user/pass")
	if err != nil {
		fmt.Println("Error opening databse:", err)
		return
	}

	// Register database with mysqldump
	dumper, err := mysqldump.Register(db, "dumps", time.ANSIC)
	if err != nil {
		fmt.Println("Error registering databse:", err)
		return
	}

	// Dump database to file
	_, err = dumper.Dump()
	if err != nil {
		fmt.Println("Error dumping:", err)
		return
	}

	// Close dumper and connected database
	dumper.Close()
}
