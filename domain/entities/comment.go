package entities

type Comment struct {
	GormModel
	Message string `gorm:"notNull"`
	UserID  uint   `gorm:"notNull;index"`
	PhotoID uint   `gorm:"notNull;index"`
}
