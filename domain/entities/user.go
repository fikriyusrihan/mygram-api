package entities

import "my-gram/domain/dto"

type User struct {
	GormModel
	Username string `gorm:"notNull;uniqueIndex"`
	Password string `gorm:"notNull"`
	Email    string `gorm:"notNull;uniqueIndex"`
	Age      uint   `gorm:"notNull"`
}

func (u *User) FromRequest(request *dto.UserRequest) {
	u.Username = request.Username
	u.Password = request.Password
	u.Email = request.Email
	u.Age = request.Age
}

func (u *User) ToResponse() *dto.UserResponse {
	return &dto.UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Age:       u.Age,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}
