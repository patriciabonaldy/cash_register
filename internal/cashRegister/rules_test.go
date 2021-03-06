package cashRegister

import (
	"encoding/json"
	"testing"

	"github.com/patriciabonaldy/cash_register/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_getRules(t *testing.T) {
	tests := []struct {
		name    string
		request models.Item
		want    []Rule
	}{
		{
			name: "unknown product",
			request: models.Item{
				Product: models.Product{
					Code:  "UNKNOWN",
					Name:  "UNKNOWN",
					Price: 5,
				},
				Quantity: 1,
				Total:    5,
			},
			want: []Rule{},
		},
		{
			name: "buy_two_by_one_free",
			request: models.Item{
				Product: models.Product{
					Code: "VOUCHER",
					Name: "VOUCHER",
				},
				Quantity: 2,
			},
			want: []Rule{
				{
					Name:     "buy_two_by_one_free",
					Desc:     "A 2-for-1 special on VOUCHER items.",
					Product:  "VOUCHER",
					Quantity: 2,
					NewPrice: 0,
					fn:       discountBuyingTwoGetOneFree,
				},
			},
		},
		{
			name: "buy_three_or_more_new_price",
			request: models.Item{
				Product: models.Product{
					Code: "TSHIRT",
					Name: "TSHIRT",
				},
				Quantity: 3,
			},
			want: []Rule{
				{
					Name:     "buy_three_or_more_new_price",
					Desc:     "If you buy 3 or more, the price per unit should be 19.00€.",
					Product:  "TSHIRT",
					Quantity: 3,
					NewPrice: 19,
					fn:       discount3OrMore,
				},
			},
		},
	}

	err := LoadRulesConfig()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RulesEngine(tt.request)
			want, err := json.Marshal(tt.want)
			require.NoError(t, err)

			got, err := json.Marshal(r)
			assert.Equal(t, string(want), string(got))
		})
	}
}
