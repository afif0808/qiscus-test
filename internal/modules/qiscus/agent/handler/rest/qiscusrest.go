package rest

import (
	"context"

	"github.com/afif0808/qiscus-test/internal/payloads"
	"github.com/labstack/echo/v4"
)

type usecase interface {
	AllocateAgent(context.Context, payloads.QiscusAgentAllocation) error
}

type AgentRestHandler struct {
	uc usecase
}

func NewAgentRestHandler(uc usecase) AgentRestHandler {
	return AgentRestHandler{uc: uc}
}

func (arh AgentRestHandler) Mount(root *echo.Group) {
	g := root.Group("/qiscus/webhook/agent/")
	g.POST("allocate", arh.allocateAgent)
}

func (arh AgentRestHandler) allocateAgent(c echo.Context) error {
	var payload payloads.QiscusAgentAllocation
	err := c.Bind(&payload)
	if err != nil {

	}
	ctx := c.Request().Context()
	err = arh.uc.AllocateAgent(ctx, payload)
	if err != nil {

	}

	return nil
}
