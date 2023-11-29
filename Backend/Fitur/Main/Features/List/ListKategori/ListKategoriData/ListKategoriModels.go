package listkategoridata

type RequestKategoriList struct {
	ID       int    `gorm:"type:int(255)" json:"id"`
	Kategori string `gorm:"type:varchar(191)" json:"kategori"`
}
