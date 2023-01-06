package main

import (
	_ "erply/docs"
	"erply/middlewares"
	"log"
)

//	@title			Erply cache server
//	@version		1.0
//	@description	It reads and writes customer data using Erply API. It sues cache with Redis.

//	@contact.name	Daichi Ando
//	@contact.url	https://github.com/canigetyourwhatwhat/cloud_base_customer_management/blob/main/README.md
//	@contact.email	daichiando98@gmail.com

//	@host		localhost:9000
//	@schemes	http
func main() {
	r := middlewares.NewRouter()
	log.Fatal(r.Run(":9000"))
}
