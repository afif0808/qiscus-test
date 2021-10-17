package payloads

type activeRoomService struct {
	RoomID string `json:"room_id"`
}

type ResolveActiveRoom struct {
	Service activeRoomService `json:"service"`
}
