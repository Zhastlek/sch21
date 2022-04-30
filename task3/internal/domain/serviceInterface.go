package domain

import "task3/internal/model"

type Service interface {
	DiscountAB(products *model.Basket, amount int)
	DiscountDE(products *model.Basket, amount int)
	DiscountEFG(products *model.Basket, amount int)
	DiscountAklm(products *model.Basket, amount int)
	DiscountTotalFive(products *model.Basket, amount int)
	DiscountTotalTen(products *model.Basket, amount int)
	DiscountTotalTwelve(products *model.Basket, amount int)
	Total(products *model.Basket)
}
