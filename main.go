package main

import (
	"log"
	"net/http"
	"os"

	agemtmodule "github.com/afif0808/qiscus-test/internal/modules/qiscus/agent"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{Message: "Hello World"})
	})
	agemtmodule.InjectAgentModule(e)

	port := os.Getenv("PORT")
	e.Start(":" + port)

}
