package wish

import (
	"github.com/devfeel/dotweb"
	"restfulapi/models"
	"restfulapi/controllers/auth"
	"restfulapi/common/utils"
	"time"
)

//获取用户待办清单
func MyWish(req dotweb.Context) error {
	res := models.Res{Code: 1}

	list, _ := models.GetWishListByUserID(auth.UserInfo.Id)
	res.Data = list
	userinfo := auth.UserInfo
	userinfo.Id = ""
	userinfo.User_login = ""
	userinfo.User_pwd = ""
	userinfo.User_status = ""

	res.Other = userinfo
	res.Code = 0

	return req.WriteJson(res)
}

//添加用户待完成清单
func AddWish(req dotweb.Context) error {
	res := models.Res{Code: 1}

	text := req.PostFormValue("text")
	if len(text) > 0 {
		err := models.Insert(&models.Wish{Id: utils.GetUUID(), User_id: auth.UserInfo.Id, Content: text, Post_date: time.Now()})
		if err == nil {
			res.Code = 0
			res.Msg = "添加成功"
		}
	}
	return req.WriteJson(res)
}
