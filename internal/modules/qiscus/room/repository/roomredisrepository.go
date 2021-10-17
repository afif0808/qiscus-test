package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/afif0808/qiscus-test/internal/domains"
	"github.com/gomodule/redigo/redis"
)

type ActiveRoomRedisRepository struct {
	rp *redis.Pool
}

func NewActiveRoomRedisRepository(rp *redis.Pool) ActiveRoomRedisRepository {
	return ActiveRoomRedisRepository{rp: rp}
}

func (repo *ActiveRoomRedisRepository) AddActiveRoom(ctx context.Context, qar domains.QiscusActiveRoom) error {
	conn := repo.rp.Get()
	buff, err := json.Marshal(qar)
	if err != nil {
		return err
	}
	_, err = conn.Do("RPUSH", "room:active", buff)
	return err
}

func (repo *ActiveRoomRedisRepository) GetAllActiveRooms(ctx context.Context) ([]domains.QiscusActiveRoom, error) {
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
		qars = append(qars, qar)
	}

	return qars, nil
}

func (repo *ActiveRoomRedisRepository) GetAgentActiveRooms(ctx context.Context, agentID int64) ([]domains.QiscusActiveRoom, error) {
	allQars, err := repo.GetAllActiveRooms(ctx)
	if err != nil {
		return nil, err
	}
	var qars []domains.QiscusActiveRoom
	for _, v := range allQars {
		if v.AgentID == agentID {
			qars = append(qars, v)
		}
	}

	return qars, nil
}

func (repo *ActiveRoomRedisRepository) GetActiveRoom(ctx context.Context, roomID string) (domains.QiscusActiveRoom, error) {
	var qar domains.QiscusActiveRoom
	allQars, err := repo.GetAllActiveRooms(ctx)
	if err != nil {
		return qar, err
	}
	for _, v := range allQars {
		if v.RoomID == roomID {
			qar = v
			break
		}
	}

	return qar, nil
}

func (repo *ActiveRoomRedisRepository) RemoveActiveRoom(ctx context.Context, roomID string) error {
	allQars, err := repo.GetAllActiveRooms(ctx)
	if err != nil {
		return err
	}
	for _, v := range allQars {
		if v.RoomID != roomID {
			continue
		}
		buff, err := json.Marshal(v)
		if err != nil {
			return err
		}
		repo.rp.Get().Do("LREM", 0, buff)
		break
	}

	return nil
}

type RoomQueueRedisRepository struct {
	rp *redis.Pool
}

func NewRoomQueueRedisRepository(rp *redis.Pool) RoomQueueRedisRepository {
	return RoomQueueRedisRepository{rp: rp}
}
func (repo *RoomQueueRedisRepository) EnqueueRoom(ctx context.Context, roomID string) error {
	conn := repo.rp.Get()
	_, err := conn.Do("RPUSH", "room:queue", roomID)
	return err
}
func (repo *RoomQueueRedisRepository) DequeueRoom(ctx context.Context) (roomID int64, err error) {
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
