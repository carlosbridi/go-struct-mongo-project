package documents

import (
	"main/src/main/domains"
	"time"
)

type InvoiceDocument struct {
	Id         string
	Customer   CustomerDocument
	Value      float64
	OrderDate  time.Time
	CreateDate time.Time
}

func (thisInvoice *InvoiceDocument) ToDomain() domains.Invoice {
	return domains.Invoice{
		Id:         thisInvoice.Id,
		Customer:   thisInvoice.Customer.ToDomain(),
		Value:      thisInvoice.Value,
		OrderDate:  thisInvoice.OrderDate,
		CreateDate: thisInvoice.CreateDate,
	}
}
