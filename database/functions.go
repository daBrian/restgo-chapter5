package database

import "chapter5/domain"

func (r *Repository) CreateInvoice(invoice domain.Invoice) domain.Invoice {
	invoice.Id = r.nextId()
	invoice.Status = "open"
	invoice.Bookings = []domain.Booking{}
	r.invoices[invoice.Id] = invoice
	return invoice
}

func (r Repository) nextId() int {
	id := 1
	for k, _ := range r.invoices {
		if k >= id {
			id = k + 1
		}
	}
	return id
}
