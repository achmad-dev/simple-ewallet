package pkg

import "testing"

func TestRandomizeStockPrice(t *testing.T) {
	// Test the RandomizeStockPrice function
	price := RandomizeStockPrice()
	if price < 100 || price > 500 {
		t.Errorf("Expected price between 100 and 500, got %f", price)
	}
}
