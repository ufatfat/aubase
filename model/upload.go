package model

type CreateWork struct {
	Class				uint8		`json:"class"`
	WorkID				uint32		`json:"work_id,omitempty" gorm:"primary_key"`
	ActivityID			uint32		`json:"activity_id,omitempty"`
	GroupID				uint32		`json:"group_id,omitempty"`
	WorkGroup			string		`json:"work_group,omitempty"`
	WorkName			string		`json:"work_name"`
	SeqID				string		`json:"seq_id"`
	LeaderName			string		`json:"leader_name"`
	LeaderOrg			string		`json:"leader_org"`
	Designers			string		`json:"designers"`
	Teacher				string		`json:"teacher,omitempty"`
	Phone				string		`json:"phone"`
	Email				string		`json:"email"`
}

type WorkGroup struct {
	GroupID				uint32		`json:"group_id"`
	GroupName			string		`json:"group_name"`
	GroupDesc			string		`json:"group_desc"`
}

type ImageInfo struct {
	WorkID				uint32		`json:"work_id"`
	ImageUrl			string		`json:"image_url"`
	ImageIndex			uint8		`json:"image_index"`
}

type FileInfo struct {
	ActivityID			string
	WorkID				string
	ImageName			string
}