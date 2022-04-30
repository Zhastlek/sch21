package adapters

import (
	"github.com/Zhastlek/school21/internal/domain"
	"github.com/Zhastlek/school21/internal/model"
)

type Calculator interface {
	CalculateTheDiscount(b *model.Basket) *model.Basket
}

type total struct {
	Total domain.DiscountServiceInterface
}

func NewTotal(totalService domain.DiscountServiceInterface) Calculator {
	return &total{
		Total: totalService,
	}
}

func (t *total) CalculateTheDiscount(b *model.Basket) *model.Basket {
	b.LinkAB = t.findLinkAB(b)
	t.Total.DiscountAB(b, b.LinkAB)

	b.LinkDE = t.findLinkDE(b)
	t.Total.DiscountDE(b, b.LinkDE)

	b.LinkEFG = t.findLinkEFG(b)
	t.Total.DiscountEFG(b, b.LinkEFG)

	b.LinkAklm = t.findLinkAklm(b)
	t.Total.DiscountAklm(b, b.LinkAklm)

	amountProducts := t.findBalanceBasket(b)
	if amountProducts >= 5 {
		t.Total.DiscountTotalTwelve(b, amountProducts)
	} else if amountProducts == 4 {
		t.Total.DiscountTotalTen(b, amountProducts)
	} else if amountProducts == 3 {
		t.Total.DiscountTotalFive(b, amountProducts)
	}
	t.Total.Total(b)
	return b
}
