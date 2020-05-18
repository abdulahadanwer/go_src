package magazine

// Subscriber holds subscribers info
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

// Employee holds employee data
type Employee struct {
	Name   string
	Salary float64
	Address
}

// Address holds addresses of different types of structs
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
