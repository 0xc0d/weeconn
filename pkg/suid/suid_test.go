package suid_test

import (
	"github.com/0xc0d/weeconn/pkg/suid"
	"testing"
	"time"
)
import "github.com/stretchr/testify/require"

func TestNewUID(t *testing.T) {
	var uid []byte
	for i := 0; i < 100; i ++ {
		uid = suid.NewUID()
		require.Equal(t, len(uid), 7,
			"short unique ID length should be 7 got %s", len(uid))
		time.Sleep(suid.TimeUnit)
	}
}

func BenchmarkNewUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		suid.NewUID()
	}
}