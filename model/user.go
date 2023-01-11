package model

type User struct {
	ID         int `gorm:"primaryKey"`
	UserName   string
	Password   string
	FirstName  string
	LastName   string
	Gender     bool
	Updated_at int64
	Created_at int64
	Disabled   bool `gorm:"default:false"`
}
