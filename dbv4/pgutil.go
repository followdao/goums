package dbv4

import (
	"regexp"
	"time"

	"github.com/jackc/pgtype"

	"emperror.dev/errors"

	"github.com/tsingson/goums/pkg/vtils"
)

// DateTimeFormat "2006-01-02 15:04:05"
const DateTimeFormat = "2006-01-02"

// ParseDay parse YYYY-MM-DD
func ParseDay(str string) (string, error) {
	regx := regexp.MustCompile(`^([\d]{4})-([\d]{2})-([\d]{2})`)
	p := regx.FindStringSubmatch(str)

	if len(p) == 4 {
		return vtils.StrBuilder(p[1], "-", p[2], "-", p[3]), nil
	}

	return "", errors.New("ParseDay ------ error")
}

// ToPgDate to pgx date
func ToPgDate(date string) (pgDate pgtype.Date, err error) {
	if len(date) == 0 {
		err = errors.New("date string is empty")
		return
	}
	var t time.Time
	t, err = time.Parse(DateTimeFormat, date)
	if err != nil {
		return
	}
	year := t.Year()
	month := t.Month()
	day := t.Day()

	pgDate = pgtype.Date{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC), Status: pgtype.Present}

	return
}

// PgDateToString to string
func PgDateToString(pgDate pgtype.Date) string {
	return pgDate.Time.Format(DateTimeFormat)
}
