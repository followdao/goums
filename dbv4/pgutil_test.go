package dbv4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPgDate(t *testing.T) {
	rawDate := "2019-10-11"
	p, err := ToPgDate(rawDate)
	assert.NoError(t, err)
	if err == nil {
		raw := PgDateToString(p)

		assert.Equal(t, raw, rawDate)

	}
}
