package liftclient

import (
	"context"
	"fmt"
	"time"

	lift "github.com/liftbridge-io/go-liftbridge"
	proto "github.com/liftbridge-io/liftbridge-api/go"

	"github.com/tsingson/goums/grpc/config/liftconfig"
)

// LiftClient new
type LiftClient struct {
	client lift.Client
	cfg    *liftconfig.LiftConfig
}

// NewLiftClient new
func NewLiftClient(ctx context.Context, cfg *liftconfig.LiftConfig) (*LiftClient, error) {
	client, err := lift.Connect(cfg.AddressList)
	if err != nil {
		return nil, err
	}

	c := &LiftClient{
		client: client,
		cfg:    cfg,
	}

	err = client.CreateStream(
		ctx, cfg.Subject, cfg.Name,
		// lift.MaxReplication(),
		//  lift.Partitions(5),
	)
	if err != nil {
		if err == lift.ErrStreamExists {
			return c, nil
		}
		return nil, err
	}

	return c, nil
}

// Stream  setup stream
func (l *LiftClient) Stream(ctx context.Context) error {
	err := l.client.CreateStream(
		ctx, l.cfg.Subject, l.cfg.Name,
		// lift.MaxReplication(),
		//  lift.Partitions(5),
	)
	if err != nil {
		if err == lift.ErrStreamExists {
			return nil
		}
		return err
	}
	return nil
}

// Publish  publish message
func (l *LiftClient) Publish(ctx context.Context, payload []byte) error {
	_, err := l.client.Publish(ctx, l.cfg.Subject,
		payload,
		//   lift.Key(keys[rand.Intn(len(keys))]),
		// lift.PartitionByKey(),
		// lift.AckPolicyLeader(),
		// lift.AckPolicyAll(),
	)
	return err
}

// Handler func(msg *proto.Message, err error)
var defaultHandler = func(msg *proto.Message, err error) {
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Unix(0, msg.Timestamp), msg.Offset, string(msg.Key), string(msg.Value))
}

// Subscribe subscribe
func (l *LiftClient) Subscribe(ctx context.Context, handler lift.Handler) {
	// lift.StartAtEarliestReceived()

	err := l.client.Subscribe(ctx, l.cfg.Name, handler, lift.StartAtTimeDelta(time.Duration(24)*time.Hour))
	if err != nil {
		//  panic(err)
	}
}

// Setup setup
func Setup(addr []string, subject string) (lift.Client, error) {
	client, err := lift.Connect(addr)
	if err != nil {
		return nil, err
	}

	err = client.CreateStream(
		context.Background(), subject, subject,
	//	lift.MaxReplication(),
	// lift.Partitions(5),
	)
	if err != nil {
		if err != lift.ErrStreamExists {
			return nil, err
		}
	}
	// fmt.Println("created stream:  ", subject)

	return client, nil
}
