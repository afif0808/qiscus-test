package rest

import (
	"context"
	"encoding/json"

	"github.com/afif0808/qiscus-test/internal/payloads"
	"github.com/labstack/echo/v4"
)

type usecase interface {
	ResolveRoom(ctx context.Context, p payloads.ResolveRoom) error
	SetResolveRoomWebhookURL(ctx context.Context, url string) error
}

type RoomRestHandler struct {
	uc usecase
}

func NewRoomRestHandler(uc usecase) RoomRestHandler {
	handler := RoomRestHandler{uc: uc}
	return handler
}

func (rrh RoomRestHandler) Mount(root *echo.Group) {
	g := root.Group("/room/room/active/")
	g.POST("resolve", rrh.resolveRoom)
}

func (rrh RoomRestHandler) resolveRoom(c echo.Context) error {
	var payload payloads.ResolveRoom
	body := c.Request().Body
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {

	}
	ctx := c.Request().Context()
	err = rrh.uc.ResolveRoom(ctx, payload)
	if err != nil {

	}
	return nil
}
