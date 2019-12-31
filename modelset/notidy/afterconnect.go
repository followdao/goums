package notidy

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const (
	SqlListenTerminal = "sqlListenTerminal"
)

func afterConnect(ctx context.Context, conn *pgx.Conn) (err error) {
	prepare := map[string]string{
		SqlListenTerminal: `listen terninal_notify`,
	}

	for key, value := range prepare {
		_, err = conn.Prepare(ctx, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
