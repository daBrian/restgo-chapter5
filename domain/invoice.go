package domain

type Invoice struct {
	Id         int    `json:"id"`
	CustomerId int    `json:"customerId"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	Status     string `json:"status"`
	Bookings   []Booking
}
type Booking struct {
	Id         int     `json:"id"`
	Day        int     `json:"day"`
	Hours      float32 `json:"hours"`
	InvoiceId  int     `json:"invoiceId"`
	ProjectId  int     `json:"projectId"`
	ActivityId int     `json:"activityId"`
}
