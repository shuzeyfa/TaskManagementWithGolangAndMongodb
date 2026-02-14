package main

import (
	"taskmanagement/router"
)

func main() {

	router := router.Router()
	router.Run(":3030")
}
