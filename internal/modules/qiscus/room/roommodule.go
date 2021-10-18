package room

import (
	agentusecase "github.com/afif0808/qiscus-test/internal/modules/qiscus/agent/usecase"
	roomrepository "github.com/afif0808/qiscus-test/internal/modules/qiscus/room/repository"
	roomusecase "github.com/afif0808/qiscus-test/internal/modules/qiscus/room/usecase"

	"github.com/labstack/echo/v4"
)

func InjectRoomModule(e *echo.Echo) {
	repo := struct {
		roomrepository.RoomSQLRepository
	}{
		RoomSQLRepository: roomrepository.NewRoomSQLRepository(nil, nil),
	}

	var usecases struct {
		roomusecase.RoomUsecase
		agentusecase.AgentUsecase
	}
	usecases.AgentUsecase = agentusecase.NewAgentUsecase(&repo)
	usecases.RoomUsecase = roomusecase.NewRoomUsecase(&repo, &usecases.AgentUsecase)

}
