package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	order,err := NewOrder(10.0,0.5)

	assert.Nil(t,err)
	assert.Equal(t,order.Price,10.0)
	assert.Equal(t,order.Tax,0.5)
	assert.Equal(t,order.FinalPrice,10.5)
}

func TestNewOrder_With_Invalid_Price(t *testing.T) {
	order,err := NewOrder(-10.0,0.5)

	assert.NotNil(t,err)
	assert.Nil(t,order)
	assert.Error(t,err,"invalid price")
}

func TestNewOrder_With_Invalid_Tax(t *testing.T) {
	order,err := NewOrder(10.0,-0.5)

	assert.NotNil(t,err)
	assert.Nil(t,order)
	assert.Error(t,err,"invalid tax")
}