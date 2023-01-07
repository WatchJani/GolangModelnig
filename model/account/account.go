package account

import (
	"log"
	"root/model/person"
	"strconv"
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

func NewAccount(Data [][]string) map[int]*Account {
	res := map[int]*Account{}

	for index := range Data {
		balancAndID := make([]int, len(Data[0])-1)
		for row := 1; row < len(Data[0]); row++ {
			number, err := strconv.ParseInt(Data[index][row], 10, 64)

			if err != nil {
				log.Fatalln(err)
			}

			balancAndID[row-1] = int(number)
		}
		if len(balancAndID) == 7 {
			res[balancAndID[0]] = New(*person.New(balancAndID[0], Data[index][0]), balancAndID[1:])
		} else {
			return nil
		}
	}

	return res
}
