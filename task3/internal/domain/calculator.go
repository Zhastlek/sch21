package domain

import (
	"github.com/Zhastlek/school21/internal/model"
)

type service struct{}

func NewDiscountService(basket *model.Basket) DiscountServiceInterface {
	return &service{}
}

func (s *service) DiscountAB(products *model.Basket, amount int) {
	numA := 0
	numB := 0
	var sum float64
	for _, value := range products.Products {
		if value.Name == "A" && numA < amount && !value.Discount {
			sum = TotalDiscountTenProcent(value.Price)
			products.ResultSum = Result(products.ResultSum, sum)
			value.Discount = true
			numA++
		} else if value.Name == "B" && numB < amount && !value.Discount {
			sum = TotalDiscountTenProcent(value.Price)
			products.ResultSum = Result(products.ResultSum, sum)
			value.Discount = true
			numB++
		}
	}
}

func (s *service) DiscountDE(products *model.Basket, amount int) {
	numD := 0
	numE := 0
	var sum float64
	for _, value := range products.Products {
		if value.Name == "D" && numD < amount && !value.Discount {
			sum = TotalDiscountFiveProcent(value.Price)
			products.ResultSum = Result(products.ResultSum, sum)
			numD++
			value.Discount = true
		} else if value.Name == "E" && numE < amount && !value.Discount {
			sum = TotalDiscountFiveProcent(value.Price)
			products.ResultSum = Result(products.ResultSum, sum)
			numE++
			value.Discount = true
		}
	}
}

func (s *service) DiscountEFG(products *model.Basket, amount int) {
	numE := 0
	numF := 0
	numG := 0
	var sum float64
	for _, value := range products.Products {
		if value.Name == "F" && numF < amount && !value.Discount {
			sum = TotalDiscountFiveProcent(value.Price)
			products.ResultSum = Result(products.ResultSum, sum)
			value.Discount = true
			numF++
		} else if value.Name == "E" && numE < amount && !value.Discount {
			sum = TotalDiscountFiveProcent(value.Price)
			products.ResultSum = Result(products.ResultSum, sum)
			value.Discount = true
			numE++
		} else if value.Name == "G" && numG < amount && !value.Discount {
			sum = TotalDiscountFiveProcent(value.Price)
			products.ResultSum = Result(products.ResultSum, sum)
			value.Discount = true
			numG++
		}
	}
}

func (s *service) DiscountAklm(products *model.Basket, amount int) {
	var numA, numKLM int
	for _, value := range products.Products {
		if value.Name == "A" && numA < amount && !value.Discount {
			products.ResultSum = Result(products.ResultSum, value.Price)
			value.Discount = true
			numA++
		} else if (value.Name == "K" || value.Name == "L" || value.Name == "M") && numKLM < amount && !value.Discount {
			sum := TotalDiscountFiveProcent(value.Price)
			products.ResultSum = Result(products.ResultSum, sum)
			value.Discount = true
			numKLM++
		}
	}
}

func (s *service) DiscountTotalFive(products *model.Basket, amount int) {
	var countProducts int
	var nums float64
	for _, value := range products.Products {
		if value.Name != "A" && value.Name != "C" && countProducts < amount && !value.Discount {
			nums += value.Price
			value.Discount = true
			countProducts++
		}
	}
	sum := TotalDiscountFiveProcent(nums)
	products.ResultSum = Result(products.ResultSum, sum)
}

func (s *service) DiscountTotalTen(products *model.Basket, amount int) {
	var countProducts int
	var nums float64
	for _, value := range products.Products {
		if value.Name != "A" && value.Name != "C" && countProducts < amount && !value.Discount {
			nums += value.Price
			value.Discount = true
			countProducts++
		}
	}
	sum := TotalDiscountTenProcent(nums)
	products.ResultSum = Result(products.ResultSum, sum)
}

func (s *service) DiscountTotalTwelve(products *model.Basket, amount int) {
	var countProducts int
	var nums float64
	for _, value := range products.Products {
		if value.Name != "A" && value.Name != "C" && countProducts < amount && !value.Discount {
			nums += value.Price
			value.Discount = true
			countProducts++
		}
	}
	sum := TotalDiscountTwelveProcent(nums)
	products.ResultSum = Result(products.ResultSum, sum)
}

func (s *service) Total(products *model.Basket) {
	var num float64
	for _, value := range products.Products {
		if !value.Discount {
			num += value.Price
			value.Discount = true
		}
	}
	products.ResultSum = Result(products.ResultSum, num)
}

func (s *service) RemoveIndex(elements []model.Product, index int) []model.Product {
	return append(elements[:index], elements[index+1:]...)
}
