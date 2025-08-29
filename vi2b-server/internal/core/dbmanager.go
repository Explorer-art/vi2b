package core

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DBInit(dbName string) {
	var err error
	db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed init SQLite")
	}

	db.AutoMigrate(&User{})
}

func DBGetUserByIP(ip string) (*User, error) {
	var user User
	
	err := db.Where("ip = ?", ip).First(&user).Error
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

func DBCreateUser(user *User) {
	db.Create(user)
}

func DBUpdateUser(user *User, userData User) {
	db.Model(user).Updates(userData)
}

func DBDeleteUser(id uint) {
	db.Delete(&User{}, id)
}
