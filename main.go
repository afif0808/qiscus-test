package main

import (
	"log"
	"os"

	agemtmodule "github.com/afif0808/qiscus-test/internal/modules/qiscus/agent"
	roommodule "github.com/afif0808/qiscus-test/internal/modules/qiscus/room"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (readDB, writeDB *gorm.DB) {
	readDB, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Panic(err)
	}
	writeDB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Panic(err)
	}
	return
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	e := echo.New()
	readDB, writeDB := InitDB()
	agemtmodule.InjectAgentModule(e, readDB, writeDB)
	roommodule.InjectRoomModule(e, readDB, writeDB)

	port := os.Getenv("PORT")
	e.Start(":" + port)

}
