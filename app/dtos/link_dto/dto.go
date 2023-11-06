package link_dto

// todo 保存之后自动import

// 输出单个信息
type LinkDTO struct {
	LinkId      string `json:"link_id"`
	NameWithUrl string `json:"name_with_url"`
	CreatedAt   string `json:"created_at" `
}
