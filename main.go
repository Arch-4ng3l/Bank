package main

import (
	"fmt"

	api "github.com/Arch-4ng3l/Bank/Api"
	database "github.com/Arch-4ng3l/Bank/Database"
)

func main() {
	psql := database.NewPostgres(5432, "localhost", "banking")
	err := psql.Connect()
	if err != nil {
		fmt.Println(err)
	}
	api.New(":3000", psql).Run()

}
