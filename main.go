package main

import (
	"erply/infra/database"
	"erply/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	//ctx := middlewares.Prep()

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()

	//customers, err := cli.GetCustomers(ctx, filter)
	//if err != nil {
	//	fmt.Println("error ----------- 4 -----------")
	//	fmt.Println(err)
	//}
	//fmt.Println("continue ----------- 4 -----------")

	//req := map[string]string{
	//	"customerID": "12",
	//	"points":     "22",
	//}
	//_, err = cli.AddCustomerRewardPoints(ctx, req)
	//if err != nil {
	//	fmt.Println("error ----------- 5 -----------")
	//	fmt.Println(err)
	//}
	//fmt.Println("continue ----------- 5 -----------")

	//temp := map[string]string{
	//	"points": "22",
	//}

	// todo:  Still don't know why it fetches all the customers
	//cs, err := client.CustomerManager.GetCustomers(ctx, temp)
	//if err != nil {
	//	fmt.Println("error ----------- 6 -----------")
	//	fmt.Println(err)
	//}
	//fmt.Println("continue ----------- 6 -----------")
	//for _, customer := range cs {
	//	fmt.Println(customer.CustomerID)
	//}

	var err error
	var db *sqlx.DB

	if err = godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	if db, err = database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	r := middlewares.NewRouter(db)
	log.Fatal(r.Run(":9000"))
}
