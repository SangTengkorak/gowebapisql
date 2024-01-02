package models

type Peoples struct {
	ID         int    `json:"ID"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Birth_year int    `json:"birth_year"`
	Email      string `json:"email"`
}

func SamplePeoples() []Peoples {
	var people1 = []Peoples{
		{ID: 5, First_name: "Jack", Last_name: "Johnson", Birth_year: 1989, Email: "singer@zmail.com"},
	}
	return people1
}
