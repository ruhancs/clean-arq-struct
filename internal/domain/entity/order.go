package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Order struct {
	ID string
	Price float64
	Tax float64
	FinalPrice float64
}

func NewOrder(price, tax float64) (*Order, error) {
	id := uuid.New().String()
	order := &Order{
		ID: id,
		Price: price,
		Tax: tax,
	}

	err := order.isValid()
	if err != nil {
		return nil,err
	}

	order.CalculateFinalPrice()
	
	return order,nil
}

func(o *Order) isValid() error {
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func(o *Order) CalculateFinalPrice() {
	finalPrice := o.Price + o.Tax
	o.FinalPrice = finalPrice
}