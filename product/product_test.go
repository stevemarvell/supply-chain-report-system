package product_test

import (
	"supply-chain-report-system/product"
	"testing"
)

func TestProduct_CalculateChainScore(t *testing.T) {
	// Create some test products and components
	p1, _ := product.NewProduct(1, "Product 1", 10, nil)
	p2, _ := product.NewProduct(2, "Product 2", 5, nil)
	p3, _ := product.NewProduct(3, "Product 3", 20, nil)

	p1.AddComponent(p2, 2)
	p1.AddComponent(p3, 1)

	p2.AddComponent(p3, 3)

	// Test that the score is calculated correctly for all products in the chain

	expectedScore3 := 20
	expectedScore2 := 5 + (expectedScore3 * 3)
	expectedScore1 := 10 + (expectedScore2 * 2) + (expectedScore3 * 1)

	if p1.Score != expectedScore1 {
		t.Errorf("%s score is incorrect: got %v, expected %v", p1.Name, p1.Score, expectedScore1)
	}
	if p2.Score != expectedScore2 {
		t.Errorf("%s score is incorrect: got %v, expected %v", p2.Name, p2.Score, expectedScore2)
	}
	if p3.Score != expectedScore3 {
		t.Errorf("%s score is incorrect: got %v, expected %v", p3.Name, p3.Score, expectedScore3)
	}

	// Test a chain change (of which there would be many test)

	p2.UpdateManufacturingScore(15)

	expectedScore3 = 20
	expectedScore2 = 15 + (expectedScore3 * 3)
	expectedScore1 = 10 + (expectedScore2 * 2) + (expectedScore3 * 1)

	if p1.Score != expectedScore1 {
		t.Errorf("%s score is incorrect: got %v, expected %v", p1.Name, p1.Score, expectedScore1)
	}
	if p2.Score != expectedScore2 {
		t.Errorf("%s score is incorrect: got %v, expected %v", p2.Name, p2.Score, expectedScore2)
	}
	if p3.Score != expectedScore3 {
		t.Errorf("%s score is incorrect: got %v, expected %v", p3.Name, p3.Score, expectedScore3)
	}
}
