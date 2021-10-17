package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/afif0808/qiscus-test/internal/domains"
	"github.com/gomodule/redigo/redis"
)

type ActiveRoomRepository struct {
	rp *redis.Pool
}

func NewActiveRoomRepository(rp *redis.Pool) ActiveRoomRepository {
	return ActiveRoomRepository{rp: rp}
}

func (repo *ActiveRoomRepository) AddActiveRoom(ctx context.Context, qar domains.QiscusActiveRoom) error {
	conn := repo.rp.Get()
	buff, err := json.Marshal(qar)
	if err != nil {
		return err
	}
	_, err = conn.Do("RPUSH", "room:active", buff)
	return err
}
func (repo *ActiveRoomRepository) GetAgentActiveRooms(ctx context.Context, agentID int64) ([]domains.QiscusActiveRoom, error) {
	conn := repo.rp.Get()
	data, err := conn.Do("LRANGE", "room:active", 0, -1)
	if err != nil {
		return nil, err
	}

	list, isTypeCorrect := data.([]interface{})
	if !isTypeCorrect {
		return nil, errors.New("tpye is invalid")
	}

	var qars []domains.QiscusActiveRoom

	for _, v := range list {
		buff, isTypeCorrect := v.([]byte)
		if !isTypeCorrect {
			return nil, errors.New("tpye is invalid")
		}
		var qar domains.QiscusActiveRoom

		err = json.Unmarshal(buff, &qar)
		if err != nil {
			return nil, err
		}
		if qar.AgentID == agentID {
			qars = append(qars, qar)
		}
	}

	return qars, nil
}

type RoomQueueRepository struct {
	rp *redis.Pool
}

func NewRoomQueueRepository(rp *redis.Pool) RoomQueueRepository {
	return RoomQueueRepository{rp: rp}
}
func (repo *RoomQueueRepository) EnqueueRoom(ctx context.Context, roomID string) error {
	conn := repo.rp.Get()
	_, err := conn.Do("RPUSH", "room:queue", roomID)
	return err
}
func (repo *RoomQueueRepository) DequeueRoom(ctx context.Context) (roomID int64, err error) {
	conn := repo.rp.Get()
	data, err := conn.Do("LPOP", "room:queue")
	if err != nil {
		return
	}
	roomID, isTypeCorrect := data.(int64)
	if !isTypeCorrect {
		err = errors.New("type is not valid")
		return
	}
	return
}
