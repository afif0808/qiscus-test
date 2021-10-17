package usecase

import (
	"context"

	"github.com/afif0808/qiscus-test/internal/domains"
	"github.com/afif0808/qiscus-test/internal/payloads"
)

type repository interface {
	GetActiveRoom(ctx context.Context, roomID string) (domains.QiscusActiveRoom, error)
	RemoveActiveRoom(ctx context.Context, roomID string) error
	DequeueRoom(ctx context.Context) (roomID string, err error)
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

func (ruc RoomUsecase) ResolveActiveRoom(ctx context.Context, p payloads.ResolveActiveRoom) error {
	// Get active room by given room id
	// Remove active room
	// Dequeue room
	// Check if room is still waiting & unresolved
	// Assign agent
	qar, err := ruc.repo.GetActiveRoom(ctx, p.Service.RoomID)
	if err != nil {
		return err
	}
	err = ruc.repo.RemoveActiveRoom(ctx, p.Service.RoomID)
	if err != nil {
		return err
	}
	roomID, err := ruc.repo.DequeueRoom(ctx)
	if err != nil {
		return err
	}

	roomInfo, err := ruc.getRoomInformation(ctx, roomID)
	if err != nil {
		return err
	}
	if !roomInfo.IsWaiting || roomInfo.IsResolved {
		return nil
	}

	return ruc.ucs.AssignAgent(ctx, payloads.QiscusAgentAssignment{
		AgentID: qar.AgentID,
		RoomID:  roomID,
	})

}

func (ruc RoomUsecase) getRoomInformation(ctx context.Context, roomID string) (domains.QiscusRoomInformation, error) {
	return domains.QiscusRoomInformation{}, nil
}
