package entities

import "time"

type UserModels struct {
	ID        uint64     `gorm:"column:id;primaryKey" json:"id"`
	Email     string     `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Password  string     `gorm:"column:password;type:VARCHAR(255)" json:"password"`
	Role      string     `gorm:"column:role;type:VARCHAR(255)" json:"role"`
	Name      string     `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (UserModels) TableName() string {
	return "users"
}
