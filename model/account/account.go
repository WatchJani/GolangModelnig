package account

import (
	"root/model/person"
)

type Account struct {
	Person   person.Person
	Currency map[string]int
}

func New(Person person.Person, Currency []int) *Account {
	mapCurrency := make(map[string]int)

	currencyState := [...]string{"USD", "GBP", "EUR", "RUB", "JPY", "INR"}

	for index, value := range currencyState {
		mapCurrency[value] = Currency[index]
	}

	return &Account{
		Person:   Person,
		Currency: mapCurrency,
	}
}

func NewAccount(p *person.Person, c []int) map[int]*Account {
	Accounts := map[int]*Account{}

	Accounts[p.PersonID] = New(*p, c)

	return Accounts
}
