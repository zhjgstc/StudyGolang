package models

import "time"

//查询用户输入的手机和密码是否正确
func GetUserInfoModelByStructForLogin(phone string, pwd string) (Userinfo, error) {
	InitDB()
	var user Userinfo
	error := DB.Where(&Userinfo{User_login: phone, User_pwd: pwd}).First(&user).Error
	return user, error
}


//检查用户是否存在
func IsUserInfoExist(phone string) bool {
	InitDB()
	var user Userinfo
	error := DB.Where(&Userinfo{User_login: phone}).First(&user).Error
	if error == nil {
		return true
	}
	return false
}


type Userinfo struct {
	Id              string
	User_login      string
	User_pwd        string
	Display_name    string
	User_registered time.Time
	User_status     string
	User_birth      time.Time
	Total           int
}

func (Userinfo) TableName() string {
	return "user_info"
}