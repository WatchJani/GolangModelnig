package main

import (
	"fmt"
	"root/data"
	"root/model/account"
)

func main() {

	// person := person.New(1115215151, "Janko")

	account := account.NewAccount(data.DataParser())

	// fmt.Println(account[8674652543].Currency["EUR"])
	//USD", "GBP", "EUR", "RUB", "JPY", "INR
	for index, users := range account {
		fmt.Printf("%d \t %s %d %d %d %d %d %d\n", index+1, users.Person.Name, users.Currency["USD"], users.Currency["GBP"], users.Currency["EUR"], users.Currency["RUB"], users.Currency["JPY"], users.Currency["INR"])
	}
}
