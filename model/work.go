package model

type (
	WorkInfo struct {
		WorkID			uint32		`json:"work_id"`
		WorkGroup		string		`json:"work_group"`
		WorkIndex		uint32		`json:"work_index"`
		WorkImages		[]WorkImage	`json:"work_images" gorm:"-"`
	}
	WorkImage struct {
		ImageID			uint32		`json:"image_id"`
		ImageUrl		string		`json:"image_url"`
		ImageIndex		uint8		`json:"image_index"`
	}
)