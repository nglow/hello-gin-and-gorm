package dto

// BookUpdateDto is a model that client use when updating a book
type BookUpdateDto struct {
	Id uint64 `json:"id" form:"id" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
	Description string `json:"Description" form:"Description" binding:"required"`
	UserId uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

// BookCreateDto is a model that client use when create a new book
type BookCreateDto struct {
	Title string `json:"title" form:"title" binding:"required"`
	Description string `json:"Description" form:"Description" binding:"required"`
	UserId uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
