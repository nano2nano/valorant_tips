package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nano2nano/valorant_tips/internal/route"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	router := route.Init()
	print(os.Getenv("PORT"))
	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
