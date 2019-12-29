package flatums

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
		ServiceStatus:     0,
		ServiceExpiration: 0,
		SerialNumber:      "aaaa",
		ActiveCode:        "bbbb",
		AccessRole:        "tvbox",
		Operation:         "UPDATE",
	}
	b := va.Byte()

	vb := ByteTerminalProfileT(b)
	as.Equal(va.SerialNumber, vb.SerialNumber)
}
