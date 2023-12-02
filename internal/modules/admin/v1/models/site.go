package models

type Site struct {
	ID         int    `json:"id" gorm:"column:id"`
	Root       string `json:"root" gorm:"column:root"`
	SiteName   string `json:"site_name" gorm:"column:site_name"`
	IdentifyID int    `json:"identify_id" gorm:"column:identify_id"`
}