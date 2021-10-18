package repository

import (
	"context"
	"testing"

	"github.com/afif0808/qiscus-test/internal/domains"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestAddRoom(t *testing.T) {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", ":6379")
		},
	}

	repo := NewRoomRedisRepository(&pool)
	err := repo.AddRoom(context.Background(), domains.QiscusRoom{
		ID:      "100",
		AgentID: 100,
	})

	assert.NoError(t, err)

}

func TestGetRooms(t *testing.T) {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", ":6379")
		},
	}

	repo := NewRoomRedisRepository(&pool)

	qars, err := repo.GetAgentRooms(context.Background(), 99)
	assert.NoError(t, err)
	assert.Equal(t, true, len(qars) > 0)

}
