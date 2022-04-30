package main

import (
	"encoding/json"
	"fmt"
	"os"
	"task3/internal/adapters"
	"task3/internal/domain"
	"task3/internal/model"
)

func main() {
	basket := model.NewBasket()
	basketService := domain.NewService(basket)
	basketCalculation := adapters.NewTotal(basketService)

	res, _ := os.ReadFile("./input")
	json.Unmarshal(res, &basket)

	result := basketCalculation.Calculation(basket)
	fmt.Println("result sum---->", result.ResultSum)
}
