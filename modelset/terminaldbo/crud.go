package terminaldbo

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"

	"github.com/tsingson/goums/apis/flatums"
)

// Insert insert valid terminal
func (s *TerminalDbo) Insert(ctx context.Context, in *flatums.TerminalProfileT) (id int64, err error) {
	err = s.pool.QueryRow(ctx, sqlInsertTerminal, in.SerialNumber, in.ActiveCode).Scan(&id)
	return id, err
}

// InsertList insert valid terminal
func (s *TerminalDbo) InsertList(ctx context.Context, in *flatums.TerminalListT) (rows int64, err error) {
	var tx pgx.Tx
	tx, err = s.pool.Begin(ctx)
	if err != nil {
		return 0, err
	}

	for _, va := range in.List {
		// serial,trial_code,active_code,filename,apk_type,zone,access_role
		var re pgconn.CommandTag
		re, err = tx.Exec(ctx, sqlInsertTerminal, va.SerialNumber, va.ActiveCode)
		if err != nil {
			break
		}
		rows = rows + re.RowsAffected()
	}
	if err != nil {
		_ = tx.Rollback(ctx)
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		_ = tx.Rollback(ctx)
		return 0, err
	}
	return rows, nil
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
		t.Operation = flatums.NotifyTypeupdate
	}
	return t, err
}
