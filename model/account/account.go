package account

import (
	"root/model/currency"
	"root/model/person"
)

type Account struct {
	Person   person.Person
	Currency map[string]*currency.Currency
}

func New(Person person.Person, Currency currency.Currency) *Account {
	mapCurrency := map[string]*currency.Currency{}

	mapCurrency[string(Currency.Currency)] = &Currency

	return &Account{
		Person:   Person,
		Currency: mapCurrency,
	}
}

func NewAccounts(p *person.Person, c *currency.Currency) map[int]*Account {
	Accounts := map[int]*Account{}

	Accounts[p.PersonID] = New(*p, *c)

	return Accounts
}
