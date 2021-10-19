package room

import (
	agentusecase "github.com/afif0808/qiscus-test/internal/modules/qiscus/agent/usecase"
	roomuseresthandler "github.com/afif0808/qiscus-test/internal/modules/qiscus/room/handler/rest"
	roomrepository "github.com/afif0808/qiscus-test/internal/modules/qiscus/room/repository"
	roomusecase "github.com/afif0808/qiscus-test/internal/modules/qiscus/room/usecase"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func InjectRoomModule(e *echo.Echo, readDB, writeDB *gorm.DB) {
	repo := struct {
		roomrepository.RoomSQLRepository
	}{
		RoomSQLRepository: roomrepository.NewRoomSQLRepository(readDB, writeDB),
	}

	var usecases struct {
		roomusecase.RoomUsecase
		agentusecase.AgentUsecase
	}
	usecases.AgentUsecase = agentusecase.NewAgentUsecase(&repo)
	usecases.RoomUsecase = roomusecase.NewRoomUsecase(&repo, &usecases.AgentUsecase)

	restHandler := roomuseresthandler.NewRoomRestHandler(usecases.RoomUsecase)
	restHandler.Mount(e.Group(""))
}
