package rest

import (
	"context"
	"encoding/json"

	"github.com/afif0808/qiscus-test/internal/payloads"
	"github.com/labstack/echo/v4"
)

type usecase interface {
	ResolveActiveRoom(ctx context.Context, p payloads.ResolveActiveRoom) error
}

type RoomRestHandler struct {
	uc usecase
}

func NewRoomRestHandler(uc usecase) RoomRestHandler {
	return RoomRestHandler{uc: uc}
}

func (rrh RoomRestHandler) Mount(root *echo.Group) {
	g := root.Group("/room/room/active/")
	g.POST("resolve", rrh.resolveActiveRoom)
}

func (rrh RoomRestHandler) resolveActiveRoom(c echo.Context) error {
	var payload payloads.ResolveActiveRoom
	body := c.Request().Body
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {

	}
	ctx := c.Request().Context()
	err = rrh.uc.ResolveActiveRoom(ctx, payload)
	if err != nil {

	}
	return nil
}
