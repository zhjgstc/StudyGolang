package models

import "time"

type Apptoken struct {
	Id       string
	Usertype string
	Code     string
	Addtime  time.Time
	Userid   string
}

func (Apptoken) TableName() string {
	return "app_token"
}

