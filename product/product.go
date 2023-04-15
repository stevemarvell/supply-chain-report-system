package product

import "fmt"

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

func (p *Product) CalculateScore() {
	// Start with the manufacturing score
	score := p.ManufacturingScore

	// Add the scores of each component, taking quantity into account
	for _, c := range p.Components {
		score += c.Product.Score * c.Quantity
	}

	// Only update the product score if it has changed
	if p.Score != score {
		p.Score = score

		// Update the scores of all products that use this product
		for _, usedBy := range p.UsedBy {
			usedBy.CalculateScore()
		}
	}
}
