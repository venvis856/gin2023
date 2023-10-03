package jump_proxy_sdk

type SearchImgReq struct {
	Limit     int      `form:"limit" json:"limit" binding:"required"`
	ImgUrl    string   `form:"imgUrl" json:"imgUrl" binding:"required"`
	MsgId     string   `form:"msgId" json:"msgId" binding:"required"`
	DeviceNos []string `form:"deviceNos" json:"deviceNos" binding:"required"`
}

func (*SearchImgReq) api() string {
	return "/search_img"
}

type SearchImgRes struct {
	ResCommon
	Data any `json:"data"`
}

func (c *JumpProxyClient) SearchImg(req *SearchImgReq) (res SearchImgRes, err error) {
	err = c.request(req, &res)
	return
}
