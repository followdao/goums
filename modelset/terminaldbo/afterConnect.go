package terminaldbo

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const (
	sqlInsertTerminal = "sqlInsertTerminal"
	SqlListenTerminal = "sqlListenTerminal"
)

func afterConnect(ctx context.Context, conn *pgx.Conn) (err error) {
	prepare := map[string]string{
		sqlInsertTerminal: `insert into ums.terminal (serial_number,active_code) values ($1, $2) returning id`,
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
