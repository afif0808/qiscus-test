package main

import (
	"log"
	"os"

	"github.com/afif0808/qiscus-test/internal/domains"
	agemtmodule "github.com/afif0808/qiscus-test/internal/modules/qiscus/agent"
	roommodule "github.com/afif0808/qiscus-test/internal/modules/qiscus/room"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (readDB, writeDB *gorm.DB) {
	os.Setenv("DATABASE_URL", "postgres://qgfynlnxlzimzj:c329c1c3f7bda60d627599b2ed9ea57e0862b0331b75c40c56563d54c05c5387@ec2-54-164-56-10.compute-1.amazonaws.com:5432/dbqskvejr09ate")
	readDB, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Panic(err)
	}
	writeDB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Panic(err)
	}
	err = writeDB.AutoMigrate(&domains.QiscusRoom{})
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
