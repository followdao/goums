package terminaldbo

import (
	"context"

	"github.com/jackc/pgtype"

	"github.com/tsingson/goums/apis/flatums"
)

// Insert insert valid terminal
func (s *TerminalDbo) Insert(ctx context.Context, in *flatums.TerminalProfileT) (id int64, err error) {
	err = s.pool.QueryRow(ctx, sqlInsertTerminal, in.SerialNumber, in.ActiveCode).Scan(&id)
	return id, err
}

// InsertList insert valid terminal
func (s *TerminalDbo) InsertList(ctx context.Context, serial, activeCode string) (id int64, err error) {
	err = s.pool.QueryRow(ctx, sqlInsertTerminal, serial, activeCode).Scan(&id)
	return id, err
}

func (s *TerminalDbo) Update(ctx context.Context, userID int64, activeStatus bool,
	maxActiveSession int64, serviceStatus int8) (int64, error) {
	re, err := s.pool.Exec(ctx, sqlUpdateTerminal, activeStatus, maxActiveSession, serviceStatus, userID)
	if err != nil {
		return 0, err
	}
	return re.RowsAffected(), nil
}

func (s *TerminalDbo) Active(ctx context.Context, serialNumber, activeCode, apkType string) (*flatums.TerminalProfileT, error) {
	t := new(flatums.TerminalProfileT)
	var activeDate, serviceExpiration pgtype.Timestamp
	err := s.pool.QueryRow(ctx, sqlActiveTerminal, activeCode, serialNumber, apkType).Scan(&t.UserID,
		&t.SerialNumber, &t.ActiveCode, &t.ActiveStatus, &activeDate, &t.MaxActiveSession,
		&t.AccessRole, &t.ServiceStatus, &serviceExpiration)
	if err == nil {
		t.ActiveDate = activeDate.Time.Unix()
		t.ServiceExpiration = serviceExpiration.Time.Unix()
		t.Operation = "UPDATE"
	}
	return t, err
}
