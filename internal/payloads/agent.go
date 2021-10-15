package payloads

type candidateAgent struct {
}

type QiscusAgentAllocation struct {
	Candidate candidateAgent `json:"candidate_agent"`
}
