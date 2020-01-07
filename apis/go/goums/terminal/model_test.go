package terminal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTerminalProfileT_Byte(t *testing.T) {
	as := assert.New(t)

	va := &TerminalProfileT{
		UserID:            0,
		ActiveStatus:      false,
		ActiveDate:        0,
		MaxActiveSession:  0,
		ServiceStatus:     ServiceStatusTypedefault,
		ServiceExpiration: 0,
		SerialNumber:      "aaaa",
		ActiveCode:        "bbbb",
		AccessRole:        "tvbox",
		Operation:         NotifyTypeupdate,
	}
	b := va.Byte()

	vb := ByteTerminalProfileT(b)
	as.Equal(va.SerialNumber, vb.SerialNumber)
}
