package person

type Person struct {
	PersonID int
	Name     string
}

func New(ID int, Name string) *Person {
	return &Person{
		ID, Name,
	}
}
