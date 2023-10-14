package gateway

import "clean-arq-struct/internal/domain/entity"

type OrderRepositoryInterface interface {
	Create(order *entity.Order) error
}