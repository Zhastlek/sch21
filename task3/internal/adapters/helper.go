package adapters

import (
	"github.com/Zhastlek/school21/internal/model"
	"github.com/Zhastlek/school21/pkg"
)

func (t *total) findLinkAB(b *model.Basket) int {
	var numA, numB, linkAB int
	for _, value := range b.Products {
		if value.Name == "A" && !value.Discount {
			numA++
		} else if value.Name == "B" && !value.Discount {
			numB++
		}
	}
	if numA > 0 && numB > 0 {
		linkAB = pkg.Compare(numA, numB)
	}
	return linkAB
}

func (t *total) findLinkDE(b *model.Basket) int {
	var numD, numE, linkDE int
	for _, value := range b.Products {
		if value.Name == "D" && !value.Discount {
			numD++
		} else if value.Name == "E" && !value.Discount {
			numE++
		}
	}
	linkDE = pkg.Compare(numD, numE)
	return linkDE
}

func (t *total) findLinkEFG(b *model.Basket) int {
	var numE, numF, numG, linkEFG int
	for _, value := range b.Products {
		if value.Name == "E" && !value.Discount {
			numE++
		} else if value.Name == "F" && !value.Discount {
			numF++
		} else if value.Name == "G" && !value.Discount {
			numG++
		}
	}
	linkEFG = pkg.Compare(numE, numF)
	linkEFG = pkg.Compare(linkEFG, numG)
	return linkEFG
}

func (t *total) findLinkAklm(b *model.Basket) int {
	var numA, numKLM, linkAklm int
	for _, value := range b.Products {
		if value.Name == "A" && !value.Discount {
			numA++
		} else if (value.Name == "K" || value.Name == "L" || value.Name == "M") && !value.Discount {
			numKLM++
		}
	}
	linkAklm = pkg.Compare(numA, numKLM)
	return linkAklm
}

func (t *total) findBalanceBasket(b *model.Basket) int {
	countProduct := 0
	for _, value := range b.Products {
		if (value.Name != "A" && value.Name != "C") && value.Discount == false {
			countProduct++
		}
	}
	return countProduct
}
