package dto

import "dora/Backstage/model"

type UserDTO struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}

// ToUserDTO 返回数据的筛选
func ToUserDTO(user model.User) UserDTO {
	return UserDTO{
		Name: user.Name,
		Phone: user.Phone,
	}
}