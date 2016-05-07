package goperson

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) Fullname() string {
	return p.FirstName + " " + p.LastName
}
