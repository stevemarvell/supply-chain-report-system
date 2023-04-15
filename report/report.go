package report

import (
	"fmt"
	"strings"
	"supply-chain-report-system/product"
	"sync"
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
		"",
		fmt.Sprintf("Manufacturing Score: %d", p.ManufacturingScore),
		"Components:",
	}
	for _, c := range p.Components {
		lines = append(lines, fmt.Sprintf("- %d x %s: %d",
			c.Quantity,
			c.Product.Name,
			c.Product.Score*c.Quantity))
	}
	lines = append(lines, "", fmt.Sprintf("Total Score: %d", p.Score))

	return &Report{
		Product:     p,
		ReportLines: lines,
	}, nil
}

func GenerateReports(products []*product.Product) ([]*Report, error) {
	type result struct {
		report *Report
		err    error
	}
	results := make(chan result, len(products))
	wg := sync.WaitGroup{}
	for _, p := range products {
		wg.Add(1)
		go func(p *product.Product) {
			defer wg.Done()
			r, err := GenerateReport(p)
			results <- result{r, err}
		}(p)
	}
	wg.Wait()
	close(results)
	var reports []*Report
	for res := range results {
		if res.err != nil {
			return nil, fmt.Errorf("error generating report: %v", res.err)
		}
		reports = append(reports, res.report)
	}
	return reports, nil
}

func ReGenerateReports(p *product.Product) ([]*Report, error) {
	dirty := make(map[int]*product.Product)

	var makeUsersDirty func(q *product.Product)

	makeUsersDirty = func(q *product.Product) {
		for _, user := range q.UsedBy {
			_, ok := dirty[user.ID]
			if !ok {
				dirty[user.ID] = user
				makeUsersDirty(user)
			}
		}
	}

	// Mark all affected products as dirty
	dirty[p.ID] = p
	makeUsersDirty(p)

	var reports []*Report

	// Regenerate reports for all dirty products
	for _, dirtyProduct := range dirty {
		report, err := GenerateReport(dirtyProduct)
		if err != nil {
			return nil, fmt.Errorf("error generating report", err)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func (r *Report) String() string {
	return strings.Join(r.ReportLines, "\n")
}
