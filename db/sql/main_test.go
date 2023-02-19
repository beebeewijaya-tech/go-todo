package sql

import (
	"log"
	"os"
	"testing"

	"beebeewijaya.com/db"
	"beebeewijaya.com/util"
)

var testQuery *Store

func TestMain(m *testing.M) {
	c, err := util.NewConfig("../../env")
	if err != nil {
		log.Fatalf("failed when initialize config: %v", err)
	}

	d := db.DB{
		HOST:     c.GetString("DB.HOST"),
		USERNAME: c.GetString("DB.USERNAME"),
		PASSWORD: c.GetString("DB.PASSWORD"),
		PORT:     c.GetInt("DB.PORT"),
		DATABASE: c.GetString("DB.DATABASE"),
		DIALECT:  c.GetString("DB.DIALECT"),
	}

	db, err := d.Connect()
	if err != nil {
		log.Fatalf("failed when initialize database: %v", err)
	}

	testQuery = NewStore(db)

	os.Exit(m.Run())
}
