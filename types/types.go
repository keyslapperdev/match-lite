package types

const CompanyKey = "student"
const StudentKey = "company"

var StudentIDX int
var CompanyIDX int

type Student struct {
	ID   int
	Name string `json:"name" bind:"required"`
}

func NewStudent(name string) *Student {
	StudentIDX++

	return &Student{
		ID:   StudentIDX,
		Name: name,
	}
}

type Company struct {
	ID   int
	Name string `json:"name" bind:"required"`
}

func NewCompany(name string) *Company {
	CompanyIDX++

	return &Company{
		ID:   CompanyIDX,
		Name: name,
	}
}
