package models

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	SetupMigrate()
	SetupSeeder()
}

func SetupMigrate() {
	Db.AutoMigrate(&Bidang{})
	Db.AutoMigrate(&Jabatan{})
	Db.AutoMigrate(&Honorer{})
	Db.AutoMigrate(&Pegawai{})
	Db.AutoMigrate(&SuratMasuk{})
	Db.AutoMigrate(&User{})
}

func SetupSeeder() {
	hash, err := HashPassword("admin")
	if err != nil {
		panic(err)
	}

	user := User{ID: 1, Username: "admin", Password: hash, Role: "admin"}

	Db.Save(&user)

}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
