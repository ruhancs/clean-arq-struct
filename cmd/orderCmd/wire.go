//go:build wireinject
// +build wireinject

package main

import (
	"clean-arq-struct/internal/domain/gateway"
	"clean-arq-struct/internal/event"
	repository "clean-arq-struct/internal/infra/database"
	"clean-arq-struct/internal/infra/web"
	"clean-arq-struct/internal/usecase"
	"clean-arq-struct/pkg/events"
	"database/sql"

	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	repository.NewOrderRepository,
	//cada vez que encontrar um productRepositoryInterface instacia o productRepository
	wire.Bind(new(gateway.OrderRepositoryInterface), new(*repository.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	//cada vez que encontrar um productRepositoryInterface instacia o productRepository
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUsecase(db *sql.DB, evDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase{
	//criacao da injecao de dependencia
	wire.Build(
		//oque sera utilizado para criar a injecao de dependecia
		setRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)

	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

//anotations no topo, especie de tag para fazer a transformacao
//wire ira verificar as anotations