package jump_proxy_sdk

type GetCameraSnapReq struct {
	CameraNo string `form:"cameraNo" json:"cameraNo" binding:"required"`
	DeviceNo string `form:"deviceNo" json:"deviceNo" binding:"required"`
	MsgId    string `form:"msgId" json:"msgId" binding:"required"`
}

func (*GetCameraSnapReq) api() string {
	return "/get_camera_snap"
}

type GetCameraSnapRes struct {
	ResCommon
	Data any `json:"data"`
}

func (c *JumpProxyClient) GetCameraSnap(req *GetCameraSnapReq) (res GetCameraSnapRes, err error) {
	err = c.request(req, &res)
	return
}
