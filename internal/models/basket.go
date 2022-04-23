package models

const (
	Voucher                = "VOUCHER"
	Tshirt                 = "TSHIRT"
	Pants                  = "PANTS"
	VoucherQuantity        = 1
	PriceDiscount          = 19
	TshirtQuantityDiscount = 3
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

// Discount3OrMore function
// Check if client buy 3 or more the same type
// then we will apply a new price
func Discount3OrMore(item Item) Item {
	var discountAmount float64
	product := item.Product
	if item.Quantity >= TshirtQuantityDiscount {
		discountAmount = PriceDiscount * float64(item.Quantity)
		item.Total = discountAmount

		return item
	}

	discountAmount = product.Price * float64(item.Quantity)
	item.Total = discountAmount

	return item
}

// DiscountBuyingTwoGetOneFree function
// Check if client buy 1 or more the same type
// gift one free
func DiscountBuyingTwoGetOneFree(item Item) Item {
	product := item.Product
	total := product.Price * float64(item.Quantity)
	if item.Quantity >= VoucherQuantity {
		item.Quantity++
	}

	item.Total = total

	return item
}
