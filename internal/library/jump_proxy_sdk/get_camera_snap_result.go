package jump_proxy_sdk

type GetCameraSnapResultReq struct {
	MsgId string `form:"msgId" json:"msgId" binding:"required"`
}

func (*GetCameraSnapResultReq) api() string {
	return "/get_camera_snap_result"
}

type GetCameraSnapResultRes struct {
	ResCommon
	Data struct {
		CameraNo string `json:"cameraNo"`
		SnapUrl  string `json:"snapUrl"`
	} `json:"data"`
}

func (c *JumpProxyClient) GetCameraSnapResult(req *GetCameraSnapResultReq) (res GetCameraSnapResultRes, err error) {
	err = c.request(req, &res)
	return
}
