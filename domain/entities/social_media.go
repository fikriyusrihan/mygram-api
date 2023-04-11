package entities

type SocialMedia struct {
	GormModel
	Name   string `gorm:"notNull"`
	URL    string `gorm:"notNull"`
	UserID int    `gorm:"notNull;index"`
}
