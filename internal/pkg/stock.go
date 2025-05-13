package pkg

import "math/rand"

func RandomizeStockPrice() float64 {
	randomValue := float64(100 + rand.Intn(400))
	return randomValue
}
