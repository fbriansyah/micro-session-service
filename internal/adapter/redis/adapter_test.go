package redisclient

import (
	"context"
	"testing"

	"github.com/fbriansyah/micro-session-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestFoundData(t *testing.T) {
	key := uuid.New()
	data := util.RandomString(8)
	err := testAdapter.SetData(context.Background(), key.String(), data, cacheDuration)
	require.NoError(t, err)

	data2, err := testAdapter.GetData(context.Background(), key.String())
	require.NoError(t, err)

	require.Equal(t, data, data2)
}

func TestDataNotFound(t *testing.T) {
	key := uuid.New()

	_, err := testAdapter.GetData(context.Background(), key.String())
	require.Error(t, err, ErrorNotFound)
}
