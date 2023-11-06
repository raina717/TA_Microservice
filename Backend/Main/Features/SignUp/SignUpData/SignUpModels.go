package signupdata

//user input
type UserSignUp struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

//input to database
type UsersSignUpInput struct {
	FirstName   string `gorm:"type:varchar(191)" json:"first_name"`
	LastName    string `gorm:"type:varchar(191)" json:"last_name"`
	Email       string `gorm:"type:varchar(191);unique;" json:"email"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	PhoneNumber string `gorm:"type:varchar(20)" json:"phone_number"`
	CreateAt    int64  `gorm:"type:bigint" json:"create_at"`
}
