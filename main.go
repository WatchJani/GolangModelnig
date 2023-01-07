package main

import (
	"fmt"
	"root/data"
	"root/model/account"
)

func main() {

	// person := person.New(1115215151, "Janko")

	account := account.NewAccount(data.DataParser())

	fmt.Println(account[8674652543].Currency["EUR"])
}
