package model

type VotedInfo struct {
	VotedWorkIds		string	`json:"voted_work_ids"`
	Name				string	`json:"name"`
}

type WorkOrder struct {
	WorkID				uint32	`json:"work_id"`
	WorkIndex			uint16	`json:"work_index"`
	WorkName			string	`json:"work_name"`
	CurrentVotesNum		uint16	`json:"current_votes_num"`
}