package request

type UpdateBookRequest struct {
	Title string `json:"title" binding:"min=3,max=100"`
	Page  int    `json:"page" binding:"gte=1"`
}