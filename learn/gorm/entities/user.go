package entities

import "time"

type User struct {
	ID        int       `gorm:"autoIncrement;<-:create"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email;unique"`
	Password  string    `gorm:"column:password"`
	IsAdmin   bool      `gorm:"column:is_admin"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	NotField  string    `gorm:"-"`
	// Name      Name      `gorm:"embedded"`
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

func (user *User) TableName() string {
	return "users"
}
