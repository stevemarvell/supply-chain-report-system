package main

import (
	"fmt"
	"supply-chain-report-system/product"
	"supply-chain-report-system/report"
)

func main() {
	// Create some test products and components
	p1, _ := product.NewProduct(1, "Product 1", 10.0, nil)
	p2, _ := product.NewProduct(2, "Product 2", 5.0, nil)
	p3, _ := product.NewProduct(3, "Product 3", 20.0, nil)

	p1.AddComponent(p2, 3)
	p1.AddComponent(p3, 2)

	p2.AddComponent(p3, 1)

	// Generate a report for Product 3
	fmt.Println("-----------------------------------------")
	productReport3, err := report.GenerateReport(p3)
	if err != nil {
		fmt.Println("Error generating report:", err)
		return
	}
	fmt.Println(productReport3)

	// Generate a report for Product 1
	fmt.Println("-----------------------------------------")
	productReport2, err := report.GenerateReport(p2)
	if err != nil {
		fmt.Println("Error generating report:", err)
		return
	}
	fmt.Println(productReport2)

	// Generate a report for Product 1
	fmt.Println("-----------------------------------------")
	productReport1, err := report.GenerateReport(p1)
	if err != nil {
		fmt.Println("Error generating report:", err)
		return
	}
	fmt.Println(productReport1)
}
