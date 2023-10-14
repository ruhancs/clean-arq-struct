package service

import (
	"clean-arq-struct/internal/infra/grpc/pb"
	"clean-arq-struct/internal/usecase"
	"clean-arq-struct/internal/usecase/dto"
	"context"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(usecase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: usecase,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	inputUseCaseDto := dto.InputOrderDto{
		Price: float64(in.Price),
		Tax: float64(in.Tax),
	}
	outputUseCaseDto,err := o.CreateOrderUseCase.Execute(inputUseCaseDto)
	if err != nil {
		return nil,err
	}

	return &pb.CreateOrderResponse{
		Price: float32(outputUseCaseDto.Price),
		Tax: float32(outputUseCaseDto.Tax),
		FinalPrice: float32(outputUseCaseDto.FinalPrice),
	},nil
}