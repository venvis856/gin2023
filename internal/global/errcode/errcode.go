package errcode

type ErrCode int

const (
	SUCCESS ErrCode = 0 //success

	ERROR           ErrCode = -1   //server error
	TOKEN_FAIL      ErrCode = 401  //请求无效
	TOKEN_FORBIDDEN ErrCode = 403  //无本系统权限
	ERROR_PARAMS    ErrCode = 1003 // 参数错误

	ERROR_LOGIN ErrCode = 1004 // 登录失败

	ERROR_NO_LOGIN ErrCode = 1005 // 未登录

	FAILURE ErrCode = 2000 // 操作失败

	FAILURE_INSERT ErrCode = 2001 // 操作insert db失败

	FAILURE_UPDATE ErrCode = 2002 // 操作update db失败

	FAILURE_DELETE ErrCode = 2003 // 操作delete db失败

	FORBIDDEN    ErrCode = 4000 // 禁止访问
	ERROR_SERVER ErrCode = 5000 // 禁止访问

	FORBIDDEN_SN ErrCode = 4001 // 禁止SN访问

	FORBIDDEN_IP ErrCode = 4002 // 禁止IP访问

	FORBIDDEN_METHOD ErrCode = 4003 // 禁止访问方法

	FORBIDDEN_NO_PERMISSION ErrCode = 4004 // 无权限访问
	ALARM_CONFIG_TIME_ERROR ErrCode = 4005 // 时间与已有配置配置有重叠
)

type ErrWrap struct {
	errCode ErrCode
	err     error
}

func (i ErrCode) Error() string {
	return i.String()
}

func (i ErrCode) Wrap(err error) ErrWrap {
	return ErrWrap{errCode: i, err: err}
}

func (ew ErrWrap) Error() string {
	return ew.err.Error()
}

func (ew ErrWrap) ErrCode() int {
	return int(ew.errCode)
}
