package domains

import "time"

type QiscusRoom struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	AgentID   int64     `json:"agent_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type QiscusRoomInformation struct {
	IsWaiting  bool `json:"is_waiting"`
	IsResolved bool `json:"is_resolved"`
}
