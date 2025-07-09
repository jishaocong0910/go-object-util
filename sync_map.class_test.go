package o

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSyncMap(t *testing.T) {
	r := require.New(t)
	m := NewSyncMap[string, any]()
	m.Put("aaa", "111")
	r.Equal("111", m.Get("aaa"))
	m.Put("AAA", "222")
	r.Equal("111", m.Get("aaa"))
	r.Equal("222", m.Get("AAA"))

	r.Equal("bbb", m.GetIfAbsent("333", func(k string) any {
		return "bbb"
	}))
	r.Equal(int64(3), m.Len())

	r.Equal("bbb", m.GetIfAbsent("333", func(k string) any {
		return "ccc"
	}))
	r.Equal(int64(3), m.Len())
	r.True(m.ContainsKeys("aaa", "333"))
	r.False(m.ContainsKeys("aaa", "333", "x"))
	r.True(m.ContainsAnyKeys("x", "y", "aaa"))
	r.False(m.ContainsAnyKeys("x", "y", "z"))
}
