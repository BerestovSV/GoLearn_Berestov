package orderCustomer

type OrderCustomer struct {
	name     string
	surname  string
	phoneNum string
}

func New(name string, surname string, phoneNum string) OrderCustomer {
	return OrderCustomer{
		name:     name,
		surname:  surname,
		phoneNum: phoneNum,
	}
}
