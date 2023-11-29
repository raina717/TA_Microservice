package listkabupatendata

type RequestKabupatenList struct {
	ID        int    `gorm:"type:int(255)" json:"id"`
	Kabupaten string `gorm:"type:varchar(191)" json:"kabupaten"`
}
