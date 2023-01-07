package main

import (
	"fmt"
	"root/model/account"
	"root/model/person"
)

func main() {

	person := person.New(11, "Janko")

	account := account.NewAccount(person, []int{50000, 60000, 70000, 5000000, 6000000, 8000000})

	fmt.Println(account[11].Currency["EUR"])
}
