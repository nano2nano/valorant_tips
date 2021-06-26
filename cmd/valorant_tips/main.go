package main

import (
	"os"

	"github.com/nano2nano/valorant_tips/internal/route"
)

func main() {
	router := route.Init()
	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
