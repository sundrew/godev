package utools

func Max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func Average(num1 float64, num2 float64) float64 {
	return (num1 + num2) / 2.0
}
