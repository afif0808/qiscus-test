package usecase

import (
	"context"

	"github.com/afif0808/qiscus-test/internal/domains"
	"github.com/afif0808/qiscus-test/internal/payloads"
)

type repository interface {
	GetRoom(ctx context.Context, roomID string) (domains.QiscusRoom, error)
	RemoveRoom(ctx context.Context, roomID string) error
	DequeueRoom(ctx context.Context) (domains.QiscusRoom, error)
}

// Common usecases
type usecases interface {
	AssignAgent(ctx context.Context, p payloads.QiscusAgentAssignment) error
}

type RoomUsecase struct {
	repo repository
	ucs  usecases
}

func NewRoomUsecase(repo repository, ucs usecases) RoomUsecase {
	return RoomUsecase{repo: repo, ucs: ucs}
}

func (ruc RoomUsecase) ResolveRoom(ctx context.Context, p payloads.ResolveRoom) error {
	// Get active room by given room id
	// Remove active room
	// Dequeue room
	// Check if room is still waiting & unresolved
	// Assign agent
	qar, err := ruc.repo.GetRoom(ctx, p.Service.RoomID)
	if err != nil {
		return err
	}
	err = ruc.repo.RemoveRoom(ctx, p.Service.RoomID)
	if err != nil {
		return err
	}
	room, err := ruc.repo.DequeueRoom(ctx)
	if err != nil {
		return err
	}

	return ruc.ucs.AssignAgent(ctx, payloads.QiscusAgentAssignment{
		AgentID: qar.AgentID,
		RoomID:  room.ID,
	})

}
