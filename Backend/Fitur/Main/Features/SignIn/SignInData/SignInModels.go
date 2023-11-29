package signindata

// import "database/sql"

type RequestUsersLogin struct {
	Email    string `gorm:"type:varchar(191)" json:"email"`
	Password string `gorm:"type:varchar(191)" json:"password"`
}

type ResponseUsersLogin struct {
	Id          string  `gorm:"type:varchar(191)" json:"id"`
	FullName    string  `gorm:"type:varchar(191)" json:"full_name"`
	Email       string  `gorm:"type:varchar(191)" json:"email"`
	Password    string  `gorm:"type:varchar(191)" json:"password"`
	PhoneNumber int     `gorm:"type:int(20)" json:"phone_number"`
	Gender      *string `gorm:"type:varchar(191)" json:"gender"`
	Address     *string `gorm:"type:varchar(191)" json:"address"`
	Status      string  `gorm:"type:varchar(191)" json:"status"`
}
