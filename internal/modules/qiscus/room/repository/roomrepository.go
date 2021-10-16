package repository

import (
	"context"

	"github.com/afif0808/qiscus-test/internal/domains"
)

type ActiveRoomRepository struct {
}

func NewActiveRoomRepository() ActiveRoomRepository {
	return ActiveRoomRepository{}
}

func (repo *ActiveRoomRepository) AddActiveRoom(ctx context.Context, qar domains.QiscusActiveRoom) error {

	return nil
}
func (repo *ActiveRoomRepository) GetAgentActiveRooms(ctx context.Context, agentID int64) ([]domains.QiscusActiveRoom, error) {
	return nil, nil
}

type RoomQueueRepository struct {
}

func NewRoomQueueRepository() RoomQueueRepository {
	return RoomQueueRepository{}
}
func (repo *RoomQueueRepository) EnqueueRoom(ctx context.Context, roomID string) error {
	return nil
}
