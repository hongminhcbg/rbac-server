package models

type User struct {
	ID       int32  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	MetaData string `gorm:"column:meta_data"`
}

func (User) TableName() string {
	return "users"
}

type UserAndRole struct {
	ID     int32 `gorm:"column:id"`
	UserID int32 `gorm:"column:user_id"`
	RoleID int32 `gorm:"column:role_id"`
}

func (UserAndRole) TableName() string {
	return "users_roles"
}

type UserAndGroup struct {
	ID      int32 `gorm:"column:id"`
	UserID  int32 `gorm:"column:user_id"`
	GroupID int32 `gorm:"column:group_id"`
}

func (UserAndGroup) TableName() string {
	return "users_groups"
}
