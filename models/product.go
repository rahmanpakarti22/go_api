package models

type Product struct {
	ID         int64  `gorm: "primarykey" json:"id"`
	NamaProduk string `gorm:"type:varchar(200)" json:"nama_produk"`
	Deskripsi  string `gorm:"type:text" json:"deskripsi"`
}
