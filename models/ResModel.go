package models


//数据请求返回的结果
type Res struct {
	Code  int
	Msg   string
	Data  ResData
	Other ResData
}

type ResData interface {
}


