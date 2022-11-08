package user

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"username" gorm:"type:varchar(40)"`
	Password string `json:"password" gorm:"size:255"`
	Created  int64  `gorm:"autoCreateTime"` //Saniye cinsinden oluşturulma tarihi.
}

func NewUser(userName, password string) *User {
	return &User{
		UserName: userName,
		Password: password,
	}
}
