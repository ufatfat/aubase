package model

type WorkToVote struct {
	WorkInfo
	Images []WorkImages	`json:"images"`
}

type WorkImages struct {
	ImageIndex uint8	`json:"image_index"`
	ImageUrl string		`json:"image_url"`
}

type WorkInfo struct {
	ID uint64			`json:"id"`
	WorkIndex string 	`json:"work_index"`
	IsNegative bool		`json:"is_negative"`
}

type VoteInfo struct {
	WorkID		uint64	`json:"work_id"`
	Negative	bool	`json:"negative"`
}