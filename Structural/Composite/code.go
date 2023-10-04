package Composite

// component.go
type Component interface {
	CalculateValue() float64
}

// product.go
type Product struct {
	value float64
}

func (p *Product) CalculateValue() float64 {
	return p.value
}

// box.go
type Box struct {
	components []Component
	value      float64
}

func (b *Box) CalculateValue() float64 {
	result := b.value
	for _, component := range b.components {
		result += component.CalculateValue()
	}
	return result
}

func (b *Box) AddComponent(c Component) {
	b.components = append(b.components, c)
}
