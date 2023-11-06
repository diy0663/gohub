package user_dto

import "github.com/diy0663/gohub/app/models/user"

func ConvertUserToDTO(user *user.User) *UserDTO {
	if user == nil {
		return nil
	}
	return &UserDTO{
		UserId:    user.GetStringID(),
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		Name:      user.Name,
		RoleName:  user.Role.Name,
	}
}

func ConvertUsersToDTO(users []*user.User) []*UserDTO {
	if users == nil {
		return nil
	}
	var dtos []*UserDTO
	for _, user := range users {
		dtos = append(dtos, ConvertUserToDTO(user))
	}
	return dtos
}
