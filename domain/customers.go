package domain

type Customer struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Projects []Project `json:"projects"`
	UserId   string    `json:"userId"`
}
type Project struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CustomerId int    `json:"customerId"`
}
