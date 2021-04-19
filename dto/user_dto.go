package dto

// UserUpdateDto is used by client when PUT update profile
type UserUpdateDto struct {
	Id uint64 `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

// UserCreateDto is used when register a user
type UserCreateDto struct {
	Id uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
}