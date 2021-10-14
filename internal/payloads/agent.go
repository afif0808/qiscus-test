package payloads

type candidateAgent struct {
}

type AgentAllocationPayload struct {
	Candidate candidateAgent `json:"candidate_agent"`
}
