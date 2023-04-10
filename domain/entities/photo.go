package entities

type Photo struct {
	GormModel
	Title    string `gorm:"notNull"`
	Caption  string
	PhotoURL string `gorm:"notNull"`
	UserID   uint   `gorm:"notNull;index"`
}
