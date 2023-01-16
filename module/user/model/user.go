package model

type User struct {
	ID        int    `json:"id" gorm:"column:id"`
	UserName  string `json:"username" gorm:"column:username"`
	Password  string `json:"password" gorm:"column:password"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Gender    bool   `json:"gender" gorm:"column:gender"`
	// Updated_at time.Time `json:"updated_at" gorm:"column:updated_at"`
	// Created_at time.Time `json:"created_at" gorm:"column:created_at"`
	Disabled bool `json:"disabled" gorm:"default:false"`
}
type UserCreate struct {
	UserName  string `json:"username" gorm:"column:username"`
	Password  string `json:"password" gorm:"column:password"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Gender    bool   `json:"gender" gorm:"column:gender"`
	// Updated_at time.Time `json:"updated_at" gorm:"column:updated_at"`
	// Created_at time.Time `json:"created_at" gorm:"column:created_at"`
	Disabled bool `json:"disabled" gorm:"default:false"`
}
type UserUpdate struct {
	UserName  string `json:"username" gorm:"column:username"`
	Password  string `json:"password" gorm:"column:password"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Gender    bool   `json:"gender" gorm:"column:gender"`
	// Updated_at time.Time `json:"updated_at" gorm:"column:updated_at"`
	// Created_at time.Time `json:"created_at" gorm:"column:created_at"`
	Disabled bool `json:"disabled" gorm:"default:false"`
}

func (User) TableName() string { return "User" }
func (UserCreate) TableName() string {
	return User{}.TableName()
}
