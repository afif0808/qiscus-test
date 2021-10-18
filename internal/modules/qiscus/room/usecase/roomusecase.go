package usecase

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"

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

func (ruc RoomUsecase) SetResolveRoomWebhookURL(ctx context.Context, webhookURL string) error {
	c := http.Client{}
	reqURL := os.Getenv("QISCUS_API_BASE_URL") + "/api/v1/admin/service/assign_agent"
	body := url.Values{}
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Qiscus-App-Id", os.Getenv("QISCUS_APP_ID"))
	req.Header.Set("Authorization", os.Getenv("QISCUS_ADMIN_TOKEN"))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return errors.New("error happened")
	}

	return nil

}
