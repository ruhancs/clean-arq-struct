package usecase

import (
	"clean-arq-struct/internal/domain/entity"
	"clean-arq-struct/internal/domain/gateway"
	"clean-arq-struct/internal/usecase/dto"
	"clean-arq-struct/pkg/events"
)

type CreateOrderUseCase struct {
	OrderRepository gateway.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(repository gateway.OrderRepositoryInterface, evDispatcher events.EventDispatcherInterface, event events.EventInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: repository,
		OrderCreated: event,
		EventDispatcher: evDispatcher,
	}
}

func (usecase *CreateOrderUseCase) Execute(input dto.InputOrderDto) (dto.OutputOrderDto, error) {
	order, err := entity.NewOrder(input.Price, input.Tax)
	if err != nil {
		return dto.OutputOrderDto{}, err
	}

	err = usecase.OrderRepository.Create(order)
	if err != nil {
		return dto.OutputOrderDto{}, err
	}

	output := dto.OutputOrderDto{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	usecase.OrderCreated.SetPayload(output)
	usecase.EventDispatcher.Dispatch(usecase.OrderCreated)

	return output, nil
}
