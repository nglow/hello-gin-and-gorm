package entity

// Book struct represents books table in database
type Book struct {
	Id uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserId uint64 `gorm:"not null" json:"-"`
	User user `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
}
