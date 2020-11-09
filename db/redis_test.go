package db_test

import (
	"github.com/0xc0d/weeconn/db"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	require.NotNilf(t, db.NewRedisClient(), "redis client is nil")
}
