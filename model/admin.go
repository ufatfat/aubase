package model

type VotedInfo struct {
	VotedWorkIds		string
	Name				string
}

type WorkOrder struct {
	WorkID				uint32
	WorkIndex			uint16
	WorkName			string
	CurrentVotesNum		uint16
}