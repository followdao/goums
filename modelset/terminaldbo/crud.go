package terminaldbo

import (
	"context"
)

// InsertTerminal insert valid terminal
func (s *TerminalDbo) InsertTerminal(ctx context.Context, serial, activeCode string) (id int64, err error) {
	err = s.pool.QueryRow(ctx, sqlInsertTerminal, serial, activeCode).Scan(&id)
	return id, err
}

func (s *TerminalDbo) UpdateTerminal(ctx context.Context, userID int64, activeStatus bool,
	maxActiveSession int64, serviceStatus int8) (int64, error) {
	re, err := s.pool.Exec(ctx, sqlUpdateTerminal, activeStatus, maxActiveSession, serviceStatus, userID)
	if err != nil {
		return 0, err
	}
	return re.RowsAffected(), nil
}
