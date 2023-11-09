package signindata

type RequestUsersLogin struct {
	Email    string `gorm:"type:varchar(191)" json:"email"`
	Password string `gorm:"type:varchar(191)" json:"password"`
}

type ResponseUsersLogin struct {
	Id          string `gorm:"type:varchar(191)" json:"id"`
	FirstName   string `gorm:"type:varchar(191)" json:"first_name"`
	LastName    string `gorm:"type:varchar(191)" json:"last_name"`
	Email       string `gorm:"type:varchar(191)" json:"email"`
	PhoneNumber int    `gorm:"type:int(20)" json:"phone_number"`
	Gender      string `gorm:"type:varchar(191)" json:"gender"`
	Address     string `gorm:"type:varchar(191)" json:"address"`
	Status      string `gorm:"type:varchar(191)" json:"status"`
}
