package dbconfig

import (
	"github.com/sonyarouje/simdb/db"
)

func DbConfig() *db.Driver {
	driver, err := db.New("dbs")
	if err != nil {
		panic(err)
	}
	return driver
}
