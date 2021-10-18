package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/afif0808/qiscus-test/internal/domains"
	"github.com/afif0808/qiscus-test/internal/payloads"
)

type repository interface {
	AddRoom(ctx context.Context, qar domains.QiscusRoom) error
	EnqueueRoom(ctx context.Context, room domains.QiscusRoom) error
	GetAgentActiveRooms(ctx context.Context, agentID int64) ([]domains.QiscusRoom, error)
}

type AgentUsecase struct {
	repo              repository
	agentRoomCapacity int
}

func NewAgentUsecase(repo repository) AgentUsecase {
	uc := AgentUsecase{
		repo: repo,
	}
	var err error
	uc.agentRoomCapacity, err = strconv.Atoi(os.Getenv("AGENT_ROOM_CAPACITY"))
	if err != nil {
		log.Panic(err)
	}
	log.Println("Test")
	return uc
}

func (auc *AgentUsecase) AllocateAgent(ctx context.Context, p payloads.QiscusAgentAllocation) error {
	// Check if agent is available
	// If available assign to the given room
	// Otherwise add to room queue

	if p.Candidate == nil {
		auc.repo.EnqueueRoom(ctx, domains.QiscusRoom{
			ID:      p.RoomID,
			AgentID: p.Candidate.ID,
		})
		return nil
	}

	rooms, err := auc.repo.GetAgentActiveRooms(ctx, p.Candidate.ID)
	if err != nil {
		return err
	}
	isAgentRoomUnlimited := auc.agentRoomCapacity == 0
	if !isAgentRoomUnlimited && len(rooms) >= auc.agentRoomCapacity {
		auc.repo.EnqueueRoom(ctx, domains.QiscusRoom{
			ID:      p.RoomID,
			AgentID: p.Candidate.ID,
		})
		return nil
	}

	return auc.AssignAgent(ctx, payloads.QiscusAgentAssignment{
		RoomID:  p.RoomID,
		AgentID: p.Candidate.ID,
	})

}

func (auc *AgentUsecase) AssignAgent(ctx context.Context, p payloads.QiscusAgentAssignment) error {
	err := auc.assignAgent(ctx, p)
	if err != nil {
		return err
	}
	return auc.repo.AddRoom(ctx, domains.QiscusRoom{
		ID:      p.RoomID,
		AgentID: p.AgentID,
	})
}

func (auc *AgentUsecase) assignAgent(ctx context.Context, p payloads.QiscusAgentAssignment) error {
	c := http.Client{}
	body := url.Values{}
	body.Add("room_id", p.RoomID)
	body.Add("agent_id", strconv.FormatInt(p.AgentID, 10))

	url := os.Getenv("QISCUS_API_BASE_URL") + "/api/v1/admin/service/assign_agent"

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Qiscus-App-Id", os.Getenv("QISCUS_APP_ID"))
	req.Header.Set("Qiscus-Secret-Key", os.Getenv("QISCUS_SECRET_KEY"))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	mappedResp := map[string]interface{}{}

	json.NewDecoder(resp.Body).Decode(&mappedResp)
	json.NewEncoder(os.Stdout).Encode(mappedResp)

	if resp.StatusCode >= 400 {
		return errors.New("error occured")
	}

	return nil
}
