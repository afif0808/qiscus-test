package agent

import (
	"github.com/afif0808/qiscus-test/internal/modules/qiscus/agent/handler/rest"
	agentusecase "github.com/afif0808/qiscus-test/internal/modules/qiscus/agent/usecase"

	roomrepository "github.com/afif0808/qiscus-test/internal/modules/qiscus/room/repository"

	"github.com/labstack/echo/v4"
)

func InjectAgentModule(e *echo.Echo) {
	repo := struct {
		roomrepository.RoomRepository
	}{
		RoomRepository: roomrepository.NewRoomRepository(),
	}

	usecase := struct {
		agentusecase.AgentUsecase
	}{
		AgentUsecase: agentusecase.NewAgentUsecase(&repo),
	}

	restHandler := rest.NewAgentRestHandler(&usecase)
	restHandler.Mount(e.Group(""))
}