package user

import (
	"github.com/devfeel/dotweb"
	"restfulapi/models"
	"restfulapi/common/utils"
	"strconv"
	"time"
	"restfulapi/controllers/auth"
)

func UpdateBithAndYear(req dotweb.Context) error {
	res := models.Res{Code: 1}
	name := req.PostFormValue("name")
	bith := req.PostFormValue("bith")
	endyear := req.PostFormValue("endyear")

	if len(name) > 0 {
		auth.UserInfo.Display_name = name
	}
	bithday, err := time.Parse("2006-01-02 15:04:05", bith)
	if err == nil {
		auth.UserInfo.User_birth = bithday
	}
	total, err := strconv.Atoi(endyear)
	if err == nil {
		auth.UserInfo.Total = total
	}

	error := models.Update(&auth.UserInfo)
	if error == nil {
		res.Code=0
		res.Msg="保存成功"
	}else{
		res.Msg="保存失败"
	}
	return req.WriteJson(res)
}

//用户登陆
func Login(req dotweb.Context) error {
	res := models.Res{Code: 1}
	phone := req.PostFormValue("phone")
	pwd := req.PostFormValue("pwd")
	if utils.ValidatePhoneNum(phone) {
		res.Msg = "请输入正确手机号"
		return req.WriteJson(res)
	}
	if len(pwd) < 6 {
		res.Msg = "请输入密码，长度最少6个字符"
		return req.WriteJson(res)
	}
	pwd = utils.Md5(pwd)
	userInfo, err := models.GetUserInfoModelByStructForLogin(phone, pwd)

	if err != nil {
		res.Msg = "用户名或密码错误"
		return req.WriteJson(res)
	} else {
		if len(userInfo.Id) > 0 {
			models.ExecBySQL("Delete from app_token where userid='" + userInfo.Id + "'")
			token := utils.GetUUID()
			error := models.Insert(&models.Apptoken{Id: utils.GetUUID(), Usertype: "第一批用户", Code: token, Addtime: time.Now(), Userid: userInfo.Id})
			if error == nil {
				res.Code = 0
				res.Data = token
				res.Msg = "登陆成功"
			}

		}
	}

	return req.WriteJson(res)
}

//用户注册
func Register(req dotweb.Context) error {
	res := models.Res{Code: 1}
	phone := req.PostFormValue("phone")
	pwd := req.PostFormValue("pwd")
	vcode := req.PostFormValue("vcode")
	if utils.ValidatePhoneNum(phone) {
		res.Msg = "请输入正确手机号"
		return req.WriteJson(res)
	}
	if len(pwd) < 6 {
		res.Msg = "请输入密码，长度最少6个字符"
		return req.WriteJson(res)
	}
	if len(vcode) != 6 || utils.ValidateNumber6(vcode) {
		res.Msg = "请输入正确短信验证码"
		return req.WriteJson(res)
	}
	list, err := models.QueryBySQL("select * from sms_log where phone='" + phone + "' and age(now(),send_date) < '05:00:00' ORDER BY send_date desc LIMIT 1")
	if err != nil {
		res.Msg = "无效的短信验证码"
		return req.WriteJson(res)
	} else {
		for list.Next() {
			var smsModel models.Smslog
			list.Scan(&smsModel)
			code, _ := strconv.Atoi(vcode)
			if smsModel.Code != code {
				res.Msg = "请输入正确短信验证码"
				return req.WriteJson(res)
			}
		}
	}

	if !models.IsUserInfoExist(phone) {
		userModel := &models.Userinfo{Id: utils.GetUUID(), User_login: phone, User_pwd: utils.Md5(pwd), User_registered: time.Now(), User_status: "正常"}
		num, _ := strconv.Atoi(vcode)
		models.UpdateVCodeByCode(num, phone)
		error := models.Insert(userModel)
		if error == nil {
			res.Msg = "注册成功"
			res.Code = 0
		} else {
			res.Msg = "注册失败"
		}
	} else {
		res.Msg = "该手机号已存在"
	}

	return req.WriteJson(res)
}
