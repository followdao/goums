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

func (s *TerminalDbo) Active(ctx context.Context, activeCode, serialNumber, apkType string) (*terminal, error) {
	/**
	  id
	  , serial_number
	  , active_status
	  , active_date
	  , max_active_session
	  , access_role
	  , service_status
	  , service_expiration
	*/
	t := new(terminal)
	err := s.pool.QueryRow(ctx, sqlActiveTerminal, activeCode, serialNumber, apkType).Scan(&t.ID,
		&t.serial, &t.active, &t.activeDate, &t.maxActiveSession, &t.accessRole, &t.serviceStatus, &t.serviceExpiration)
	return t, err
}
