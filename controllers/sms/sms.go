package sms

import (
	"github.com/devfeel/dotweb"
	"restfulapi/models"
	"strconv"
	"restfulapi/common/utils"
	"time"
)

//发送短信
func SendACode(req dotweb.Context) error {
	res := models.Res{Code: 1}

	phone := req.PostFormValue("phone")
	smstype := req.PostFormValue("type")
	vcode, _ := strconv.Atoi(utils.GetCode())

	if utils.ValidatePhoneNum(phone) {
		res.Msg = "请输入正确手机号"
		return req.WriteJson(res)
	}

	if smstype != "注册" || smstype != "忘记密码" {
		res.Msg = "非法操作！"
		return req.WriteJson(res)
	}

	if smstype == "注册" && models.IsUserInfoExist(phone) {
		res.Msg = "该手机号已存在"
		return req.WriteJson(res)
	}

	if IsAllowSendSMS(phone, smstype, req.RemoteIP(), req.Request().UserAgent()) {
		res.Code = 0
		res.Msg = "发送成功"
		return req.WriteJson(res)
	}

	flag, _ := utils.SendSMS(phone, vcode)
	if flag {
		smsModel := &models.Smslog{Id: utils.GetUUID(), Phone: phone, Code: vcode, Send_date: time.Now(), Send_type: smstype, Status: "已发送", Ip: req.RemoteIP(), Useragent: req.Request().UserAgent()}
		err := models.Insert(smsModel)
		if err == nil {
			res.Code = 0
			res.Msg = "发送成功"
		} else {
			res.Msg = "发送失败"
		}
	} else {
		res.Msg = "发送失败"
	}

	return req.WriteJson(res)
}

//是否可以发送短信验证码
func IsAllowSendSMS(phone string, smstype string, ip string, useragent string) bool {

	listCountIp, error := models.QueryBySQL("select count(id) from sms_log where ip='" + ip + "' and useragent='" + useragent + "'")
	if error == nil {
		for listCountIp.Next() {
			var count int
			listCountIp.Scan(&count)
			if count >= 15 {
				return true
			}
		}
	}

	listCount, error := models.QueryBySQL("select count(id) from sms_log where phone='" + phone + "'")
	if error == nil {
		for listCount.Next() {
			var count int
			listCount.Scan(&count)
			if count >= 5 {
				return true
			}
		}
	}

	list, err := models.QueryBySQL("select * from sms_log where phone='" + phone + "' age(now(),send_date) < '05:00:00' ORDER BY send_date desc LIMIT 1")
	if err == nil {
		for list.Next() {
			var smsModel models.Smslog
			list.Scan(&smsModel)
			if smsModel.Send_type == smstype {
				return true
			}
		}
	}
	return false
}
