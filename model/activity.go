package model

type (
	TurnInfo struct {
		TurnID			uint32		`json:"turn_id"`
		TurnIndex		uint8		`json:"turn_index"`
		IsPositive		bool		`json:"is_positive"`
		VotesNum		uint16		`json:"votes_num"`
	}
)