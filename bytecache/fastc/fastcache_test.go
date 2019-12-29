package fastc

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/**
func TestMain(m *testing.M) {
    fmt.Println("Test Main")
    os.Exit(m.Serve())
}

*/

func TestNewFastCache(t *testing.T) {
	var c *VictoriaCache
	var err error
	k, v := []byte("key"), []byte("value1111")
	t.Run("new", func(t *testing.T) {
		c = NewVictoriaCache(WithCacheSize(1024 * 1024 * 512))
		assert.NoError(t, err)
	})
	t.Run("set ", func(t *testing.T) {
		check := c.Set(k, v)
		assert.True(t, check)
	})
	t.Run("get ", func(t *testing.T) {
		vv, check := c.Get(k)
		assert.True(t, check)
		if check {
			assert.Equal(t, vv, v)
		}
	})
	t.Run("del  ", func(t *testing.T) {
		c.Del(k)
		time.Sleep(time.Duration(10) * time.Microsecond)
		_, check := c.Get(k)
		assert.False(t, check)
	})
}

func BenchmarkFastCache_Get(b *testing.B) {
	var c *VictoriaCache

	k, v := []byte("key"), []byte("valuevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevaluevalue")

	c = NewVictoriaCache()

	_ = c.Set(k, v)

	// b.SetParallelism(128)
	b.SetParallelism(runtime.NumCPU() * 2)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = c.Get(k)
		}
	})
}
