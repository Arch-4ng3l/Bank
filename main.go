package main

import (
	"fmt"
	"log"

	api "github.com/Arch-4ng3l/Bank/Api"
	database "github.com/Arch-4ng3l/Bank/Database"
	web "github.com/Arch-4ng3l/Bank/Web"
)

func main() {
	fmt.Println("[*] STARTING SERVER [*]")
	psql := database.NewPostgres(5432, "localhost", "banking")
	err := psql.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	web.AddFrontend()
	api.New(":3000", psql).Run()

}
