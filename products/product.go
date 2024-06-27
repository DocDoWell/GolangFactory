package products

import "time"

type Product struct{
	ProductName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//example of factory patterm
func (p *Product) New() *Product{
	product := Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &product
}