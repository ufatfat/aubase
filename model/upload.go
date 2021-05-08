package model

type CreateWork struct {
	WorkID				uint32		`json:"work_id,omitempty" gorm:"primary_key"`
	GroupID				uint32		`json:"group_id"`
	WorkName			string		`json:"work_name"`
	SeqID				string		`json:"seq_id"`
	FirstDesignerOrg	string		`json:"first_designer_org"`
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