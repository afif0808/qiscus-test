package rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/afif0808/qiscus-test/internal/payloads"
	"github.com/afif0808/qiscus-test/wrapper"
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
	err := json.NewDecoder(c.Request().Body).Decode(&payload)

	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusInternalServerError, "failed", err).JSON(c.Response())
	}
	log.Println(payload)

	ctx := c.Request().Context()
	err = arh.uc.AllocateAgent(ctx, payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusInternalServerError, "failed", err).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Agent allocated").JSON(c.Response())
}
