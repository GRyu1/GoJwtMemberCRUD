package main

import (
	"github.com/gin-contrib/cors"
	"goLangJwtPrac/db"
	"goLangJwtPrac/routes"
	"log"
)

func main() {
	err := db.InitMongoDB()
	if err != nil {
		log.Fatalln("디비 연결 실패", err)
	}
	router := routes.SetupRouter()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	err = router.Run(":7070")
	if err != nil {
		log.Fatalln("서버 오픈 실패", err)
	}
}
