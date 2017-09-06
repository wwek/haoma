package api

type DefaultApi struct {
	ErrCode int         `json:"err_code"` //错误代码
	ErrMsg  string      `json:"err_msg"`  //错误信息
	Data    interface{} `json:"data"`     //数据
}
