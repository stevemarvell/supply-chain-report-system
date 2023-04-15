package product

import (
	"fmt"
	"log"
)

type Product struct {
	ID                 int
	Name               string
	ManufacturingScore int
	Components         map[int]*Component
	Score              int
	UsedBy             []*Product
}

func NewProduct(id int, name string, manufacturingScore int, components map[*Product]int) (*Product, error) {
	if name == "" {
		return nil, fmt.Errorf("product name cannot be empty")
	}
	if manufacturingScore < 0 {
		return nil, fmt.Errorf("manufacturing score cannot be negative")
	}

	p := &Product{
		ID:                 id,
		Name:               name,
		ManufacturingScore: manufacturingScore,
		Components:         make(map[int]*Component),
		Score:              manufacturingScore,
	}

	// Add the components to the product
	for c, quantity := range components {
		p.AddComponent(c, quantity)
	}

	return p, nil
}

func (p *Product) UpdateManufacturingScore(manufacturingScore int) {
	p.ManufacturingScore = manufacturingScore
	p.CalculateScore()
}

func (p *Product) AddComponent(c *Product, quantity int) {
	p.Components[c.ID] = &Component{
		Product:  c,
		Quantity: quantity,
	}
	c.UsedBy = append(c.UsedBy, p)
	p.CalculateScore()
}

func (p *Product) RemoveComponent(c *Product) {
	delete(p.Components, c.ID)
	for j, usedBy := range c.UsedBy {
		if usedBy == p {
			c.UsedBy = append(c.UsedBy[:j], c.UsedBy[j+1:]...)
			break
		}
	}
	p.CalculateScore()
}

// TopologicalSort A function to perform a topological sort on a list of products
func TopologicalSort(products []*Product) ([]*Product, error) {
	sorted := make([]*Product, 0, len(products))
	visited := make(map[int]bool)
	var visit func(p *Product)

	visit = func(p *Product) {
		visited[p.ID] = true
		for _, usedBy := range p.UsedBy {
			if !visited[usedBy.ID] {
				visit(usedBy)
			}
		}
		sorted = append(sorted, p)
	}

	for _, p := range products {
		if !visited[p.ID] {
			visit(p)
		}
	}

	// Reverse the order of the sorted list
	for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
		sorted[i], sorted[j] = sorted[j], sorted[i]
	}

	return sorted, nil
}

func (p *Product) CalculateScore() {

	dirty := make(map[int]*Product)

	var makeUsersDirty func(q *Product)

	makeUsersDirty = func(q *Product) {
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
	var products []*Product
	for _, p := range dirty {
		products = append(products, p)
	}

	// Sort the products in topological order
	sortedProducts, err := TopologicalSort(products)

	if err != nil {
		log.Printf("error calculating product scores: %v", err)
		return
	}

	// Calculate the scores of each product in the correct order
	for _, product := range sortedProducts {
		product.Score = product.ManufacturingScore
		for _, c := range product.Components {
			product.Score += c.Product.Score * c.Quantity
		}
	}
}
