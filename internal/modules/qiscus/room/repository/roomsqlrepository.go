package repository

import (
	"context"

	"github.com/afif0808/qiscus-test/internal/domains"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoomSQLRepository struct {
	readDB, writeDB *gorm.DB
}

func NewRoomSQLRepository(readDB, writeDB *gorm.DB) RoomSQLRepository {
	return RoomSQLRepository{readDB: readDB, writeDB: writeDB}
}

func (repo *RoomSQLRepository) AddRoom(ctx context.Context, room domains.QiscusRoom) error {
	room.IsActive = true
	return repo.writeDB.Save(room).Error
}

func (repo *RoomSQLRepository) GetAgentRooms(ctx context.Context, agentID int64) ([]domains.QiscusRoom, error) {
	var rooms []domains.QiscusRoom
	err := repo.readDB.Where("is_active = ?", true).Find(&rooms).Error
	return rooms, err
}

func (repo *RoomSQLRepository) GetRoom(ctx context.Context, roomID string) (domains.QiscusRoom, error) {
	var room domains.QiscusRoom
	err := repo.readDB.First(&room, roomID).Error
	return room, err
}

func (repo *RoomSQLRepository) RemoveRoom(ctx context.Context, roomID string) error {
	err := repo.writeDB.Delete(domains.QiscusRoom{}, roomID).Error
	return err
}
func (repo *RoomSQLRepository) EnqueueRoom(ctx context.Context, room domains.QiscusRoom) error {
	room.IsActive = false
	return repo.writeDB.Save(room).Error
}
func (repo *RoomSQLRepository) DequeueRoom(ctx context.Context) (domains.QiscusRoom, error) {
	var room domains.QiscusRoom
	err := repo.readDB.
		Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: false}).
		Where("is_active = ?", false).
		First(&room).Error
	if err != nil {
		return room, err
	}
	room.IsActive = true
	err = repo.writeDB.Save(&room).Error
	return room, err
}
