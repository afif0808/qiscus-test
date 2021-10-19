package payloads

type candidateAgent struct {
	ID int64 `json:"id"`
}

type QiscusAgentAllocation struct {
	RoomID    string          `json:"room_id"`
	Candidate *candidateAgent `json:"candidate_agent"`
}

type QiscusAgentAssignment struct {
	RoomID  string
	AgentID int64
}
