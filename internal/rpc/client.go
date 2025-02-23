package rpc

import (
	"github.com/osamikoyo/IM-auth/internal/config"
	"github.com/osamikoyo/IM-auth/pkg/loger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RpcClient struct{
	channel *amqp.Channel
	loger loger.Logger
	cfg *config.Config
}

func (r *RpcClient) New(cfg *config.Config) (*RpcClient, error) {
	conn, err := amqp.Dial(cfg.AmqpConnect)
	if err != nil{
		return nil, err
	}
	defer conn.Close()

	ch,err := conn.Channel()
	if err != nil{
		return nil, err
	}
	defer ch.Close()

	return &RpcClient{
		channel: ch,
		loger: loger.New(),
		cfg: cfg,
	}, nil
}

func (r *RpcClient) Send(UserID uint64) {

}