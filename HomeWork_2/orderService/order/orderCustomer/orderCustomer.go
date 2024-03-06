package orderCustomer

type OrderCustomer struct {
	Name     string
	Surname  string
	PhoneNum string
}

func New(name string, surname string, phoneNum string) OrderCustomer {
	return OrderCustomer{
		Name:     name,
		Surname:  surname,
		PhoneNum: phoneNum,
	}
}
