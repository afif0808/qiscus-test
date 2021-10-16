package repository

import (
	"context"

	"github.com/afif0808/qiscus-test/internal/domains"
)

type RoomRepository struct {
}

func NewRoomRepository() RoomRepository {
	return RoomRepository{}
}

func (rr *RoomRepository) AddActiveRoom(ctx context.Context, qar domains.QiscusActiveRoom) error {
	return nil
}
func (rr *RoomRepository) EnqueueRoom(ctx context.Context, roomID string) error {
	return nil
}
func (rr *RoomRepository) GetAgentActiveRooms(ctx context.Context, agentID int64) ([]domains.QiscusActiveRoom, error) {
	return nil, nil
}
