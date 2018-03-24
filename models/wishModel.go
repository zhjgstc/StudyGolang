package models

import "time"

type Wish struct {
	Id          string
	User_id     string
	Content     string
	Post_date   time.Time
	Finish_date time.Time
	Cancel_date time.Time
}

//获取愿望列表
func GetWishListByUserID(userid string) ([]Wish,error){
	InitDB()
	var list []Wish
	err := DB.Where(&Wish{User_id:userid}).Order("post_date desc").Find(&list).Error
	return list,err
}
