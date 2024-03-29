package main

import (
	"os"
	"unjuk_prakerja/configs"
	"unjuk_prakerja/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	e.Start(":" + os.Getenv("PORT"))

}
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env files")
	}
}
