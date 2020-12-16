package sqlstore

import "goshop/internal/app/model"

type ProductRepository struct {
	store *Store
}

func (r *ProductRepository) GetLatestProds() ([]model.Product, error) {

}
