package main

import (
	"fmt"
	"root/model/account"
	"root/model/currency"
	"root/model/person"
)

func main() {

	person := person.New(11, "Janko")
	EUR := currency.New("EUR", 1000)

	accounts := account.NewAccounts(person, EUR)

	fmt.Println(accounts[11].Currency["EUR"].Balance)
}
