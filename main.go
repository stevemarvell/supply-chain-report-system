package main

import (
	"fmt"
	"supply-chain-report-system/product"
	"supply-chain-report-system/report"
)

func main() {
	// Create some test products and components
	p1, _ := product.NewProduct(1, "Product 1", 10, nil)
	p2, _ := product.NewProduct(2, "Product 2", 5, nil)
	p3, _ := product.NewProduct(3, "Product 3", 20, nil)

	p1.AddComponent(p2, 3)
	p1.AddComponent(p3, 2)

	p2.AddComponent(p3, 1)

	products := []*product.Product{p1, p2, p3}

	for _, p := range products {
		r, err := report.GenerateReport(p)
		if err != nil {
			fmt.Println("Error generating r:", err)
			return
		}

		fmt.Println("----------------------")
		fmt.Println(r)
	}

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++")

	p2.UpdateManufacturingScore(15)

	reports, _ := report.ReGenerateReports(p2)

	for _, r := range reports {
		fmt.Println("----------------------")
		fmt.Println(r)
	}
}
