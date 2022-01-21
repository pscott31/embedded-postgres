package main

import (
	"log"
	"time"

	epg "github.com/fergusstrange/embedded-postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"vega.spike/entities"
)

func main() {
	pg := epg.NewDatabase(epg.DefaultConfig().
		Username("beer").
		Password("wine").
		Database("gin").
		Version(epg.V14).
		RuntimePath("/home/scotty/work/epg/runtime").
		Port(9876).
		StartTimeout(45 * time.Second).
		Logger(log.Writer()))

	pg.Start()

	db, err := gorm.Open(postgres.Open("postgres://beer:wine@localhost:9876/gin"))
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entities.Transfer{})
	db.AutoMigrate(&entities.Account{})
	db.AutoMigrate(&entities.Asset{})
	db.AutoMigrate(&entities.Party{})

	log.Println("yay")

}
