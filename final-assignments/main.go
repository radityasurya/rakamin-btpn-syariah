package main

import (
	"github.com/radityasurya/final-assignment/database"
	"github.com/radityasurya/final-assignment/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
