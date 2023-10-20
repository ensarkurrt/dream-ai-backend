package dao

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique_index; not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
