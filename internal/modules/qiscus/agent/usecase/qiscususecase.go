package usecase

import (
	"context"

	"github.com/afif0808/qiscus-test/internal/payloads"
)

type roomRepository interface {
	AddRoom(ctx context.Context)
	GetAgentRooms()
}

type customerQueueRepository interface {
	DequeueCustomer(ctx context.Context)
	EnqueueCustomer(ctx context.Context)
}

type repository interface {
	roomRepository
	customerQueueRepository
}

type AgentUsecase struct {
	repo              repository
	allocationChannel chan struct {
		ctx context.Context
		p   payloads.AgentAllocationPayload
	}
}

func NewAgentUsecase(repo repository) AgentUsecase {
	uc := AgentUsecase{
		repo: repo,
		allocationChannel: make(chan struct {
			ctx context.Context
			p   payloads.AgentAllocationPayload
		}),
	}
	go func() { uc.allocateAgent() }()

	return uc
}

func (auc AgentUsecase) AllocateAgent(ctx context.Context, p payloads.AgentAllocationPayload) error {
	auc.allocationChannel <- struct {
		ctx context.Context
		p   payloads.AgentAllocationPayload
	}{
		ctx: ctx, p: p,
	}
	return nil
}

func (auc AgentUsecase) allocateAgent() error {

	return nil
}
