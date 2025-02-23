package rpc

import (
	"github.com/osamikoyo/IM-auth/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RpcClient struct{
	channel *amqp.Channel
	cfg *config.Config
}