package main

import (
	"erply/middlewares"
	"log"
)

func main() {

	r := middlewares.NewRouter()
	log.Fatal(r.Run(":9000"))
}
