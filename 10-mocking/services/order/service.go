package order

import (
	"context"
	"errors"
	"github.com/vageeshabr/go-intro/10-mocking/services"
)

type Order struct {
	//store stores.OrderStorer
	smsSender   services.SmsSender
	emailSender services.EmailSender
}

func New(smsSender services.SmsSender, emailSender services.EmailSender) *Order {
	return &Order{
		smsSender:   smsSender,
		emailSender: emailSender,
	}
}

type Customer struct {
	Id    int
	Name  string
	Phone string
	Email string
}

type OrderCreate struct {
	Customer  *Customer
	ItemCount int
}

func (o *Order) Create(ctx context.Context, req *OrderCreate) error {

	if req.Customer == nil {
		return errors.New("customer details are missing")
	}

	if req.Customer.Phone == "" && req.Customer.Email == "" {
		return errors.New("phone and email are missing")
	}

	// create an order & save it to db.

	if req.Customer.Phone != "" {
		if err := o.smsSender.Send(req.Customer.Phone, "Your order is successfully placed"); err != nil {
			//
			return errors.New("internal server error")
		}
	}

	if req.Customer.Email != "" {
		if err := o.emailSender.Send(req.Customer.Email, "Your order is successfully placed", "here are the details"); err != nil {
			return errors.New("internal server error")
		}
	}

	return nil
}
