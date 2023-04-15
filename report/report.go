package report

import (
	"fmt"
	"strings"
	"supply-chain-report-system/product"
)

type Report struct {
	Product     *product.Product
	ReportLines []string
}

func GenerateReport(p *product.Product) (*Report, error) {
	if p == nil {
		return nil, fmt.Errorf("product cannot be nil")
	}

	// Generate the report lines
	lines := []string{
		fmt.Sprintf("Product: %s", p.Name),
		fmt.Sprintf("Manufacturing Score: %d", p.ManufacturingScore),
		"",
		"Components:",
	}
	for _, c := range p.Components {
		lines = append(lines, fmt.Sprintf("- %d x %s: %d",
			c.Quantity,
			c.Product.Name,
			c.Product.Score*int(c.Quantity)))
	}
	lines = append(lines, "", fmt.Sprintf("Total Score: %d", p.Score))

	return &Report{
		Product:     p,
		ReportLines: lines,
	}, nil
}

func (r *Report) String() string {
	return strings.Join(r.ReportLines, "\n")
}
