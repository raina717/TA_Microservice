package signupdata

//user input
type UserSignUp struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

//input to database
type UsersSignUpInput struct {
	FullName    string `gorm:"type:varchar(255)" json:"full_name"`
	Email       string `gorm:"type:varchar(255);unique;" json:"email"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	PhoneNumber string `gorm:"type:int(20)" json:"phone_number"`
	// CreateAt    int64  `gorm:"type:date" json:"create_at"`
}
