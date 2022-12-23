package documents

import "main/src/main/domains"

type CustomerDocument struct {
	Id      string
	Name    string
	Address string
	Active  bool
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
