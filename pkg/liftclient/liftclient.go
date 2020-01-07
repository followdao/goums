package liftclient

import (
	"context"

	lift "github.com/liftbridge-io/go-liftbridge"
)

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
