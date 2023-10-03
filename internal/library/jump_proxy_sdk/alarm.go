package jump_proxy_sdk

type AlarmReq struct {
	ImgUrl    string   `form:"imgUrl" json:"imgUrl" binding:"required"`
	CameraNo  string   `form:"cameraNo" json:"cameraNo" binding:"required"`
	StartTime string   `form:"startTime" json:"startTime" binding:"required"`
	EndTime   string   `form:"endTime" json:"endTime" binding:"required"`
	Point     []int32  `form:"point" json:"point" binding:"required"`
	AlarmType int32    `form:"alarmtype" json:"alarmtype" binding:"required"` //1 酒店的移动侦测 2 人员聚集 3 电子围栏
	Threshold int32    `form:"threshold" json:"threshold" binding:"required"`
	DeviceNo  []string `form:"deviceNo" json:"deviceNo" binding:"required"`
	MsgId     string   `form:"msgId" json:"msgId" binding:"required"`
}

func (*AlarmReq) api() string {
	return "/alarm"
}

type AlarmRes struct {
	ResCommon
	Data any `json:"data"`
}

func (c *JumpProxyClient) Alarm(req *AlarmReq) (res AlarmRes, err error) {
	err = c.request(req, &res)
	return
}
