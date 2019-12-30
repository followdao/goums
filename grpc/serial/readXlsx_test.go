package serial

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStbList(t *testing.T) {
	as := assert.New(t)

	l, err := ReadList("testdata/sn.xlsx")
	as.NoError(err)
	if err == nil {
		for k, v := range l.List {
			fmt.Printf(" %02d  \n ", k)
			fmt.Println("sn: ", v.SerialNumber, "  code: ", v.ActiveCode)
		}
	}
}
