package main

import (
	"fmt"
	"log"

	"beebeewijaya.com/api"
	"beebeewijaya.com/db"
	"beebeewijaya.com/db/sql"
	"beebeewijaya.com/util"
)

func main() {
	c, err := util.NewConfig("./env")
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

	store := sql.NewStore(db)
	s := api.NewServer(store, c)

	address := fmt.Sprintf("%s:%d", c.GetString("SERVER.ADDRESS"), c.GetInt("SERVER.PORT"))
	err = s.Start(address)
	if err != nil {
		log.Fatalf("error when starting server: %v", err)
	}
}
