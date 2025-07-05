package dto

type PostCreateDto struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostUpdateDto struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostListDto struct {
	ID     uint   `json:"id" binding:"required"`
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
}

type PostGetOrDelDto struct {
	ID uint `json:"id" binding:"required"`
}
