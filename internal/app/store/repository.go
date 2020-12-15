package store

import "goshop/internal/app/model"

type ProductRepository interface {
	GetLatestProds() ([]model.Product, error)
}
