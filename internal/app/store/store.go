package store

type Store interface {
	Product() ProductRepository
}
