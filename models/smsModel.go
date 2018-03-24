package models

import "time"

//更新验证码
func UpdateVCodeByCode(vcode int, phone string) error {
	InitDB()
	var smslog Smslog
	error := DB.Where(&Smslog{Phone: phone, Code: vcode}).First(&smslog).Error
	if error == nil {
		smslog.Status = "已使用"
		return Update(&smslog)
	}
	return error
}
type Smslog struct {
	Id        string
	Phone     string
	Code      int
	Send_date time.Time
	Send_type string
	Status    string
	Ip string
	Useragent string
}

func (Smslog) TableName() string {
	return "sms_log"
}
