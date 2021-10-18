package payloads

type RoomService struct {
	RoomID string `json:"room_id"`
}

type ResolveRoom struct {
	Service RoomService `json:"service"`
}
