package agent

import (
	"github.com/afif0808/qiscus-test/internal/modules/qiscus/agent/handler/rest"
	agentusecase "github.com/afif0808/qiscus-test/internal/modules/qiscus/agent/usecase"
	"gorm.io/gorm"

	roomrepository "github.com/afif0808/qiscus-test/internal/modules/qiscus/room/repository"

	"github.com/labstack/echo/v4"
)

func InjectAgentModule(e *echo.Echo, readDB, writeDB *gorm.DB) {
	repo := struct {
		roomrepository.RoomSQLRepository
	}{
		RoomSQLRepository: roomrepository.NewRoomSQLRepository(readDB, writeDB),
	}

	usecase := struct {
		agentusecase.AgentUsecase
	}{
		AgentUsecase: agentusecase.NewAgentUsecase(&repo),
	}

	restHandler := rest.NewAgentRestHandler(&usecase)
	restHandler.Mount(e.Group(""))
}
