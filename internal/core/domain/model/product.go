package domain

import (
	"fmt"

	domain_builder "github.com/alfariiizi/vandor/internal/core/domain/builder"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
)

type Product struct {
	*db.Product
	client *db.Client
}

func NewProductDomain(client *db.Client) domain_builder.Domain[*db.Product, *Product] {
	return domain_builder.NewDomain(
		func(e *db.Product, c *db.Client) *Product {
			return &Product{
				Product: e,
				client:  c,
			}
		}, client)
}

// TODO: Add your domain methods here
// Example methods:

func (product *Product) String() string {
	return fmt.Sprintf("Product{ID: %s}", product.ID)
}

// Add more business logic methods as needed
