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
	WorkIndex uint16 	`json:"work_index"`
}

type VoteInfo struct {
	WorkID		uint64	`json:"work_id"`
	Positive	bool	`json:"positive"`
}