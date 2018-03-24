package LifeLong

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Smslog struct{
	Id string
	User_login string
	User_pwd string
	Display_name string
	User_registered time.Time
	User_status string
	User_birth time.Time

}
func (Smslog) TableName() string{
	return "sms_log"
}

type Userinfo struct{
	Id string
	User_login string
	User_pwd string
	Display_name string
	User_registered time.Time
	User_status string
	User_birth time.Time
	Total int

}
func (Userinfo) TableName() string{
	return "user_info"
}

type Wish struct {
	Id string
	User_id string
	Content string
	post_date time.Time
	finish_date time.Time
	cancel_date time.Time
}

var DB *gorm.DB

//初始化数据库
func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=LifeLong sslmode=disable password=postgres")
	if err == nil {
		DB = db
		db.SingularTable(true)
		db.AutoMigrate(&Userinfo{}, &Smslog{}, &Wish{})
		return db, err
	}
	return nil, err
}

// 创建用户
func (user *Userinfo) InsertUserInfo() error {
	return DB.Create(user).Error
}
func (user *Userinfo) Delete()error{
	return DB.Delete(user).Error
}
func (user *Userinfo) Update() error{
	return DB.Update(user).Error
}
func(user *Userinfo) GetList() error{
	return DB.Find(user).Error
}


