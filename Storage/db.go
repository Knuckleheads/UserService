package Storage

import (
	"log"
	"rateService/prisma/db"
)

type Dbinstance struct {
	Db *db.PrismaClient
}

var DB Dbinstance

func ConnectDb() {
	log.Println("connectDB...")
	db := db.NewClient()

	if err := db.Prisma.Connect(); err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	log.Println("connected")
	DB = Dbinstance{
		Db: db,
	}
}
