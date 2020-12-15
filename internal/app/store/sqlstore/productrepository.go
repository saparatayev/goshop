package sqlstore

import "goshop/internal/app/model"

type ProductRepository struct {
	store *Store
}

func GetLatestProds() ([]model.Product, error) {

}
