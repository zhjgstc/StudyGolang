package utils

import (
	"time"
	"fmt"
	"math/rand"
	"github.com/satori/go.uuid"
	"restfulapi/common/subsms"
	"strconv"
	"encoding/json"
	"regexp"
	"crypto/md5"
	"encoding/hex"
)

//获取六位数随机码
func GetCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

//获取uuid
func GetUUID() string {
	u1, _ := uuid.NewV4()
	return u1.String()
}

//发送短信验证码
func SendSMS(phone string, code int) (bool, string) {
	messageconfig := make(map[string]string)
	messageconfig["appid"] = ""//短信密钥
	messageconfig["appkey"] = ""//短信密钥
	messageconfig["signtype"] = "md5"

	messagexsend := subsms.CreateMessageXSend()
	subsms.MessageXSendAddTo(messagexsend, phone)
	subsms.MessageXSendSetProject(messagexsend, "WmDHu3")
	subsms.MessageXSendAddVar(messagexsend, "code", strconv.Itoa(code))
	subsms.MessageXSendAddVar(messagexsend, "time", "5")
	res := subsms.MessageXSendRun(subsms.MessageXSendBuildRequest(messagexsend), messageconfig)
	result := &smsResult{}
	err := json.Unmarshal([]byte(res), result)
	if err == nil {
		if result.Status == "success" {
			return true, "发送成功"
		}
	}
	return false, "发送失败"
}

//短信回传数据类型
type smsResult struct {
	Status                    string `json:"status"`
	Send_id                   string `json:"send_id"`
	Fee                       int    `json:"fee"`
	Sms_credits               string `json:"sms_credits"`
	Transactional_sms_credits string `json:"transactional_sms_credits"`
}

//验证手机号
func ValidatePhoneNum(mobileNum string) bool {
	regular := "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

//验证是否六位数字
func ValidateNumber6(text string) bool {
	regular := "/^\\d{6}$/"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(text)
}

//md5加密
func Md5(text string) string {
	h := md5.New()
	h.Write([]byte(text)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
