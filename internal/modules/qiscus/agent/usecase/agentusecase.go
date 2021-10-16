package usecase

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/afif0808/qiscus-test/internal/domains"
	"github.com/afif0808/qiscus-test/internal/payloads"
)

type repository interface {
	AddActiveRoom(ctx context.Context, qar domains.QiscusActiveRoom) error
	EnqueueRoom(ctx context.Context, roomID int64) error
	GetAgentActiveRooms(ctx context.Context, agentID int64) ([]domains.QiscusActiveRoom, error)
}

type AgentUsecase struct {
	repo repository
}

func NewAgentUsecase(repo repository) AgentUsecase {
	uc := AgentUsecase{
		repo: repo,
	}

	return uc
}

func (auc *AgentUsecase) AllocateAgent(ctx context.Context, p payloads.QiscusAgentAllocation) error {
	// Check if agent is available
	// If available assign to the given room
	// Otherwise add to room queue

	rooms, err := auc.repo.GetAgentActiveRooms(ctx, p.Candidate.ID)
	if err != nil {
		return err
	}

	if len(rooms) >= 2 {
		auc.repo.EnqueueRoom(ctx, p.RoomID)
		return nil
	}
	err = auc.assignAgent(ctx, payloads.QiscusAgentAssignment{
		RoomID:  p.RoomID,
		AgentID: p.Candidate.ID,
	})

	if err != nil {
		return err
	}

	return auc.repo.AddActiveRoom(ctx, domains.QiscusActiveRoom{
		AgentID: p.Candidate.ID,
		RoomID:  p.RoomID,
	})
}

func (auc *AgentUsecase) assignAgent(ctx context.Context, p payloads.QiscusAgentAssignment) error {
	c := http.Client{}
	body := url.Values{}
	body.Add("room_id", strconv.FormatInt(p.RoomID, 10))
	body.Add("agent_id", strconv.FormatInt(p.AgentID, 10))

	url := os.Getenv("QISCUS_API_BASE_URL") + "/api/v1/admin/service/assign_agent"

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Qiscus-App-Id", os.Getenv("QISCUS_APP_ID"))
	req.Header.Set("Qiscus-Secret-Key", os.Getenv("QISCUS_SECREY_KEY"))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return errors.New("error occured")
	}

	return nil
}
