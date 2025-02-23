package rpc

import (
	"context"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/osamikoyo/IM-auth/internal/config"
	"github.com/osamikoyo/IM-auth/internal/data/models"
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

func (r *RpcClient) Send(UserID uint64) error {
	q, err := r.channel.QueueDeclare(
		r.cfg.RpcQueName,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil{
		return err
	}

	msgs,err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		return err
	}

	corr := randomString(32)

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	respBody, err := sonic.Marshal(&models.Req{
		ID: UserID,
	})

	err = r.channel.PublishWithContext(ctx,
		"",
		r.cfg.RpcQueName,
		false,
		false,
		amqp.Publishing{
			CorrelationId: corr,
			ContentType: "application/json",
			ReplyTo: q.Name,
			Body: respBody,
		},
	)

	var resp models.Resp

	for d := range msgs{
		if err = sonic.Unmarshal(d.Body, &resp);err != nil{
			r.loger.Error().Err(err)
			return err
		}

		if resp.Status != http.StatusOK{
			r.loger.Error().Msg(resp.Error)
		}
	}

	return nil
}