package documents

import "go_sample/src/main/domains"

type CustomerDocument struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"adress"`
	Active  bool   `json:"active"`
}

func ToDocument(customer domains.Customer) CustomerDocument {
	return CustomerDocument{
		Id:      customer.Id,
		Name:    customer.Name,
		Address: customer.Address,
		Active:  customer.Active,
	}
}

func (customer *CustomerDocument) ToDomain() domains.Customer {
	return domains.Customer{
		Id:      customer.Id,
		Name:    customer.Name,
		Address: customer.Address,
		Active:  customer.Active,
	}
}
