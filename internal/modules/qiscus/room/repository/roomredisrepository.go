package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/afif0808/qiscus-test/internal/domains"
	"github.com/gomodule/redigo/redis"
)

type RoomRedisRepository struct {
	rp *redis.Pool
}

func NewRoomRedisRepository(rp *redis.Pool) RoomRedisRepository {
	return RoomRedisRepository{rp: rp}
}

func (repo *RoomRedisRepository) AddRoom(ctx context.Context, room domains.QiscusRoom) error {
	conn := repo.rp.Get()
	buff, err := json.Marshal(room)
	if err != nil {
		return err
	}
	_, err = conn.Do("RPUSH", "room:active", buff)
	return err
}

func (repo *RoomRedisRepository) GetAllRooms(ctx context.Context) ([]domains.QiscusRoom, error) {
	conn := repo.rp.Get()
	data, err := conn.Do("LRANGE", "room:active", 0, -1)
	if err != nil {
		return nil, err
	}

	list, isTypeCorrect := data.([]interface{})
	if !isTypeCorrect {
		return nil, errors.New("tpye is invalid")
	}

	var rooms []domains.QiscusRoom

	for _, v := range list {
		buff, isTypeCorrect := v.([]byte)
		if !isTypeCorrect {
			return nil, errors.New("type is invalid")
		}
		var room domains.QiscusRoom

		err = json.Unmarshal(buff, &room)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (repo *RoomRedisRepository) GetAgentActiveRooms(ctx context.Context, agentID int64) ([]domains.QiscusRoom, error) {
	allrooms, err := repo.GetAllRooms(ctx)
	if err != nil {
		return nil, err
	}
	var rooms []domains.QiscusRoom
	for _, v := range allrooms {
		if v.AgentID == agentID {
			rooms = append(rooms, v)
		}
	}

	return rooms, nil
}

func (repo *RoomRedisRepository) GetRoom(ctx context.Context, roomID string) (domains.QiscusRoom, error) {
	var room domains.QiscusRoom
	allrooms, err := repo.GetAllRooms(ctx)
	if err != nil {
		return room, err
	}
	for _, v := range allrooms {
		if v.ID == roomID {
			room = v
			break
		}
	}

	return room, nil
}

func (repo *RoomRedisRepository) RemoveRoom(ctx context.Context, roomID string) error {
	allrooms, err := repo.GetAllRooms(ctx)
	if err != nil {
		return err
	}
	for _, v := range allrooms {
		if v.ID != roomID {
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
func (repo *RoomQueueRedisRepository) EnqueueRoom(ctx context.Context, room domains.QiscusRoom) error {
	buff, err := json.Marshal(room)
	if err != nil {
		return err
	}
	conn := repo.rp.Get()
	_, err = conn.Do("RPUSH", "room:queue", buff)
	return err
}
func (repo *RoomQueueRedisRepository) DequeueRoom(ctx context.Context) (room domains.QiscusRoom, err error) {
	conn := repo.rp.Get()
	data, err := conn.Do("LPOP", "room:queue")
	if err != nil {
		return
	}
	buff, isTypeCorrect := data.([]byte)
	if !isTypeCorrect {
		err = errors.New("type is not valid")
		return
	}
	err = json.Unmarshal(buff, &room)
	return
}
