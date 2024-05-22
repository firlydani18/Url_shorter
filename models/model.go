package models

type Link struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment;"`
	ShortURL string `json:"short_url" form:"short_url" gorm:"type: varchar(255)"`
	LongURL  string `json:"long_url" form:"long_url" gorm:"type: varchar(255)"`
}
