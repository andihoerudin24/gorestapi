package user

type UserModel struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
