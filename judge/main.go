package main

import (
	"CodeBox/repository/db"
	"CodeBox/repository/rabbitMQ"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	if _, err = os.Stat("./judge/code"); os.IsNotExist(err) {
		err = os.Mkdir("./judge/code", 0700)
		if err != nil {
			return
		}
	}

	//DB
	dbUrl := os.Getenv("DB_URL")
	//dbName := os.Getenv("DB_NAME")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//DB_URL := "postgres://" + dbName + ":" + dbPassword + "@" + dbUrl

	//Postgres db
	dialector := postgres.Open(dbUrl)
	newDb := db.NewDatabse(dialector)
	newDb.Connect()

	err = rabbitMQ.NewListenRmq()
	if err != nil {
		return
	}

}
