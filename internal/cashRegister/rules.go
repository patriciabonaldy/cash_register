package cashRegister

import "github.com/patriciabonaldy/cash_register/internal/models"

type rulesMap map[ruleName]func(request models.Item, rule Rule) func(item models.Item, rule Rule) models.Item

var _rulesMap = rulesMap{
	"buy_two_by_one_free":         buyTwoByOneFree,
	"buy_three_or_more_new_price": buyThreeOrMoreNewPrice,
}

func buyTwoByOneFree(request models.Item, rule Rule) func(item models.Item, rule Rule) models.Item {
	if request.Product.Code != rule.Product {
		return nil
	}

	if request.Quantity < rule.Quantity {
		return nil
	}

	return discountBuyingTwoGetOneFree
}

func buyThreeOrMoreNewPrice(request models.Item, rule Rule) func(item models.Item, rule Rule) models.Item {
	if request.Product.Code != rule.Product {
		return nil
	}

	if request.Quantity < rule.Quantity {
		return nil
	}

	return discount3OrMore
}

func RulesEngine(request models.Item) []Rule {
	ruleList := []Rule{}

	for name, ruleApplies := range _rulesMap {
		rConfig := configRules.Rules[name]

		if fn := ruleApplies(request, rConfig); fn != nil {
			rConfig.fn = fn
			ruleList = append(ruleList, rConfig)
		}
	}

	return ruleList
}

// discountBuyingTwoGetOneFree function
// Check if client buy 1 or more the same type
// gift one free
func discountBuyingTwoGetOneFree(item models.Item, _ Rule) models.Item {
	item.Total = item.Product.Price * float64(item.Quantity-1)

	return item
}

// discount3OrMore function
// Check if client buy 3 or more the same type
// then we will apply a new price
func discount3OrMore(item models.Item, rule Rule) models.Item {
	var discountAmount float64
	discountAmount = rule.NewPrice * float64(item.Quantity)
	item.Total = discountAmount

	return item
}
