package main

import (
	"CodeBox/repository/db"
	"CodeBox/repository/rabbitMQ"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Main.go : %v", err)
		return
	}

	dbUrl := os.Getenv("DB_URL")
	//dbName := os.Getenv("DB_NAME")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//DB_URL := "postgres://" + dbName + ":" + dbPassword + "@" + dbUrl
	//Postgres db
	dialector := postgres.Open(dbUrl)
	newDb := db.NewDatabse(dialector)
	newDb.Connect()

	//Rabbit mq emitter
	err = rabbitMQ.NewRmq()
	if err != nil {
		log.Printf("Main.go : %v", err)
		return
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	routersInit := InitRouter()
	gin.SetMode(gin.DebugMode)
	routersInit.Run(port)

}
