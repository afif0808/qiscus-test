package repository

import (
	"context"
	"testing"

	"github.com/afif0808/qiscus-test/internal/domains"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestAddActiveRoom(t *testing.T) {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", ":6379")
		},
	}

	repo := NewActiveRoomRedisRepository(&pool)
	err := repo.AddActiveRoom(context.Background(), domains.QiscusActiveRoom{
		RoomID:  "100",
		AgentID: 100,
	})

	assert.NoError(t, err)

}

func TestGetActiveRooms(t *testing.T) {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", ":6379")
		},
	}

	repo := NewActiveRoomRedisRepository(&pool)

	qars, err := repo.GetAgentActiveRooms(context.Background(), 99)
	assert.NoError(t, err)
	assert.Equal(t, true, len(qars) > 0)

}
