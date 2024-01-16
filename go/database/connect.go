package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() *gorm.DB {

	dsn := "rahma:rakesh24310@tcp(localhost:3306)/toko"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal terhubung ke database!")
	}

	db.Table(new(User).TableName()).AutoMigrate(&User{})
	db.Table(new(Bata).TableName()).AutoMigrate(&Bata{})
	db.Table(new(Roster).TableName()).AutoMigrate(&Roster{})
	db.Table(new(Listbata).TableName()).AutoMigrate(&Listbata{})
	db.Table(new(Riwayat).TableName()).AutoMigrate(&Riwayat{})
	db.Table(new(Keranjang).TableName()).AutoMigrate(&Keranjang{})

	return db
}
