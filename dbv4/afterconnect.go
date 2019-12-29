package dbv4

/**
const (
	insertRawTVODSQL = "insertRawTVODSQL"
	removeRawTVODSQL = "removeRawTVOD"
	countRawTVODSQL  = "countRawTVODSQL"
)

func afterConnect(ctx context.Context, conn *pgx.Conn) (err error) {
	//
	_, err = conn.Prepare(ctx, insertRawTVODSQL,
		`INSERT INTO data.rawepg (channel_id, channel_title, title, subtitle, date, start_time, end_time, timezone ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) `)
	if err != nil {
		return
	}
	//
	_, err = conn.Prepare(ctx, removeRawTVODSQL, `delete from data.rawepg where  date=$1 and channel_id=$2`)
	if err != nil {
		return
	}

	_, err = conn.Prepare(ctx, countRawTVODSQL,
		`select count(id) as total from data.rawepg where channel_id=$1 and date=$2`)

	return nil
}


*/
