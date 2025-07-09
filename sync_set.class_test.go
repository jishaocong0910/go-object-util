package o

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSyncSet(t *testing.T) {
	r := require.New(t)
	s := NewSyncSet[string]()
	r.True(s.Empty())
	r.Equal(int64(0), s.Len())
	s2 := NewSet[string]("a", "b", "c")
	r.False(s2.Empty())
	r.Equal(int64(3), s2.Len())

	s.Add("str")
	r.Equal(int64(1), s.Len())
	r.True(s.Contains("str"))
	s.AddSet(s2)
	r.Equal(int64(4), s.Len())
	validateElement(r, s, "a", "str", "b", "c")

	r.True(s.Contains("a", "b"))
	r.False(s.Contains("a", "b", "d"))
	r.True(s.ContainsAny("a", "d"))
	r.True(s.ContainsAny("a", "b", "d"))

	s.Remove("a")
	r.False(s.Contains("a"))
}
