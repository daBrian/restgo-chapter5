package database

import "chapter5/domain"

type Repository struct {
	invoices map[int]domain.Invoice
}

func NewRepository() Repository {
	return Repository{invoices: map[int]domain.Invoice{}}
}
