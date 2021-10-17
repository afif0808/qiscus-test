package domains

type QiscusActiveRoom struct {
	RoomID  string `json:"room_id"`
	AgentID int64  `json:"agent_id"`
}

type QiscusRoomInformation struct {
	IsWaiting  bool `json:"is_waiting"`
	IsResolved bool `json:"is_resolved"`
}
