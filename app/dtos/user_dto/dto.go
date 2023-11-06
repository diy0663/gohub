package user_dto

// 输出单个用户信息
type UserDTO struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name" `
	RoleName  string `json:"role_name"`
	CreatedAt string `json:"created_at" `
}
