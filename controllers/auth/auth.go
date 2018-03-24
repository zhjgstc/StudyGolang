package auth

import (
	"github.com/devfeel/dotweb"
	"restfulapi/models"
	"restfulapi/common/utils"
)

type AccessFmtLog struct {
	dotweb.BaseMiddlware
	exactToken string
}

func NewAccessFmtLog(index string) *AccessFmtLog {
	return &AccessFmtLog{exactToken: index}
}

var UserInfo models.Userinfo

func (m *AccessFmtLog) Handle(req dotweb.Context) error {
	res := models.Res{Code: 403}
	token := req.Request().Header.Get("token")
	phone := req.Request().Header.Get("phone")

	if len(token) <= 0 || len(phone) <= 0 || utils.ValidatePhoneNum(phone) {
		res.Msg = "登陆失效请重新登陆"
		return req.WriteJson(res)
	}

	db, err := models.InitDB()
	if err != nil {
		res.Code = 500
		res.Msg = "服务器异常"
		return req.WriteJson(res)
	}

	var apptokenModel models.Apptoken
	err = db.Where(&models.Apptoken{Code: token}).Find(&apptokenModel).Error
	if err != nil {
		res.Msg = "验证码不正确，请重新登录"
		return req.WriteJson(res)
	}

	var userModel models.Userinfo
	err = db.Where(&models.Userinfo{User_login: phone}).Find(&userModel).Error
	if err != nil {
		res.Msg = "用户信息不正确，请重新登录"
		return req.WriteJson(res)
	}

	if userModel.Id != apptokenModel.Userid {
		res.Msg = "用户信息不正确，请重新登录"
		return req.WriteJson(res)
	}
	UserInfo = userModel
	return m.Next(req)
}
