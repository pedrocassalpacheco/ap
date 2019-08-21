package types

type Employee struct {
	Id    int
	Name  string
	City string
}

func NewEmployee(id int, name string, city string) *Employee {
	return &Employee{Id: id, Name: name, City: city}
}
