package domain

func TotalDiscountFiveProcent(num float64) float64 {
	res := num - (num * 0.05)
	return res
}

func TotalDiscountTenProcent(num float64) float64 {
	res := num - (num * 0.1)
	return res
}

func TotalDiscountTwelveProcent(num float64) float64 {
	res := num - (num * 0.2)
	return res
}

func Total(nums []float64) float64 {
	var result float64
	for _, value := range nums {
		result += value
	}
	return result
}

func Result(num1, num2 float64) float64 {
	result := num1 + num2
	return result
}
