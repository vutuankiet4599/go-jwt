package request

type InsertBookRequest struct {
	Title string `json:"title" binding:"required,min=3,max=100"`
	Page  int    `json:"page" binding:"required,gte=1"`
}
