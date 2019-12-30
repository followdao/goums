package terminaldbo

import (
	"github.com/jackc/pgtype"
)

type terminal struct {
	ID int64
	serial string
	active bool
	activeDate pgtype.Date
	maxActiveSession int64
	accessRole string
	serviceStatus int32
	serviceExpiration pgtype.Timestamp
}


