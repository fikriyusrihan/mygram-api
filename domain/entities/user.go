package entities

type User struct {
	GormModel
	Username string `gorm:"notNull;uniqueIndex"`
	Password string `gorm:"notNull"`
	Email    string `gorm:"notNull;uniqueIndex"`
	Age      uint   `gorm:"notNull"`
}
