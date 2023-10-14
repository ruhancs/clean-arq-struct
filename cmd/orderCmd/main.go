package main

import (
	"clean-arq-struct/internal/event/handler"
	"clean-arq-struct/internal/infra/grpc/pb"
	"clean-arq-struct/internal/infra/grpc/service"
	"clean-arq-struct/internal/infra/web/webserver"
	"clean-arq-struct/pkg/events"
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db,err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()
	evDispatcher := events.NewEventDispatcher()
	evDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUsecase(db,evDispatcher)

	webserver := webserver.NewWebServer(os.Getenv("WEB_SERVER_PORT"))
	webOrderHandler := NewWebOrderHandler(db,evDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Server runnig on port: ", os.Getenv("WEB_SERVER_PORT"))
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", os.Getenv("GRPC_SERVER_PORT"))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_SERVER_PORT")))
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(lis)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONN"))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}