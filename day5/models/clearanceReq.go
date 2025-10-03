package models

type ClearanceReq struct {
	Discription string `json:"discription" gorm:"type:varchar(100);not null"`
	Status      string `json:"status" gorm:"type:varchar(100);not null"`
}
