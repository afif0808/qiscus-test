package payloads

type candidateAgent struct {
	ID int64 `json:"id"`
}

type QiscusAgentAllocation struct {
	RoomID    int64          `json:"room_id"`
	Candidate candidateAgent `json:"candidate_agent"`
}

type QiscusAgentAssignment struct {
	RoomID  int64
	AgentID int64
}
