package models

import (
	"reflect"
	"testing"
)

func TestDiscount3OrMore(t *testing.T) {
	t.Run("If you buy 3 or more, the price per unit should be 19.00€.",
		func(t *testing.T) {
			item := Item{
				Product: Product{
					Code:  "TSHIRT",
					Name:  "TSHIRT",
					Price: 20,
				},
				Quantity: 3,
				Total:    0,
			}
			want := Item{
				Product: Product{
					Code:  "TSHIRT",
					Name:  "TSHIRT",
					Price: 20,
				},
				Quantity: 3,
				Total:    57,
			}
			if got := Discount3OrMore(item); !reflect.DeepEqual(got, want) {
				t.Errorf("Discount3OrMore() = %v, want %v", got, want)
			}
		})
}

func TestDiscountBuyingTwoGetOneFree(t *testing.T) {
	t.Run("A 2-for-1", func(t *testing.T) {
		item := Item{
			Product: Product{
				Code:  "TSHIRT",
				Name:  "TSHIRT",
				Price: 20,
			},
			Quantity: 3,
			Total:    0,
		}
		want := Item{
			Product: Product{
				Code:  "TSHIRT",
				Name:  "TSHIRT",
				Price: 20,
			},
			Quantity: 4,
			Total:    60,
		}
		if got := DiscountBuyingTwoGetOneFree(item); !reflect.DeepEqual(got, want) {
			t.Errorf("DiscountBuyingTwoGetOneFree() = %v, want %v", got, want)
		}
	})
}
