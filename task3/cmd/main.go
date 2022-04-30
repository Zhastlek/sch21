package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Zhastlek/school21/internal/adapters"
	"github.com/Zhastlek/school21/internal/domain"
	"github.com/Zhastlek/school21/internal/model"
)

func main() {
	basket := model.NewBasket()
	basketService := domain.NewDiscountService(basket)
	basketCalculation := adapters.NewTotal(basketService)

	res, _ := os.ReadFile("./input")
	json.Unmarshal(res, &basket)

	result := basketCalculation.CalculateTheDiscount(basket)
	fmt.Println("result sum---->", result.ResultSum)
}
