package orderGood

type OrderGood struct {
	Name  string
	Price float64
	Count int
}

func New(name string, price float64, count int) OrderGood {
	return OrderGood{
		Name:  name,
		Price: price,
		Count: count,
	}
}
