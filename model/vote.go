package model

type VoteInfo struct {
	CurrentWorkID		uint32	`json:"current_work_id"`
	IsVoted				bool	`json:"is_voted"`
}

type CreateVote struct {
	CurrentWorkID		uint32
	UserID				uint32
	ActivityID			uint32
	TurnID				uint32
	VotesNum			uint16
}