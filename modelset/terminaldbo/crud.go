package terminaldbo

import (
	"context"
)

// InsertTerminal insert valid terminal
func (s *TerminalDbo) InsertTerminal(ctx context.Context, serial, activeCode string) (id int64, err error) {
	err = s.pool.QueryRow(ctx, sqlInsertTerminal, serial, activeCode).Scan(&id)
	return id, err
}
