package models

const (
	Voucher = "VOUCHER"
	Tshirt  = "TSHIRT"
	Pants   = "PANTS"
)

var (
	ProductMap = map[string]Product{
		Voucher: {Code: Voucher, Name: "Gift Card", Price: 5.00},
		Tshirt:  {Code: Tshirt, Name: "Summer T-Shirt", Price: 20.00},
		Pants:   {Code: Pants, Name: "Summer Pants ", Price: 7.50},
	}
)

type Basket struct {
	Code  string
	Items map[string]Item
	Total float64
	Close bool
}

type Product struct {
	Code  string
	Name  string
	Price float64
}

type Item struct {
	Product  Product
	Quantity int
	Total    float64
}

func NewBasket(id string) Basket {
	return Basket{
		Code:  id,
		Items: make(map[string]Item),
	}
}

func (b *Basket) CalculateTotal() {
	var total float64
	for _, i := range b.Items {
		total += i.Total
	}

	b.Total = total
}

func (i *Item) WithOutDiscount() {
	var discountAmount float64

	product := i.Product
	discountAmount = product.Price * float64(i.Quantity)
	i.Total = discountAmount
}
