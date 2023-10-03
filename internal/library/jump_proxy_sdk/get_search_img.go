package jump_proxy_sdk

type GetSearchImgReq struct {
	MsgId string `form:"msgId" json:"msgId" binding:"required"`
}

func (*GetSearchImgReq) api() string {
	return "/get_search_img"
}

type GetSearchImgRes struct {
	ResCommon
	Data map[string]struct {
		DeviceId string `json:"deviceid"`
		Result   []*struct {
			Camid  string `json:"camid"`
			Videos []*struct {
				Date       string  `json:"date"`
				Distance   float64 `json:"distance"`
				ImgUrl     string  `json:"imgurl"`
				MatchImg   string  `json:"matchimg"`
				MatchVideo string  `json:"matchvideo"`
				Seek       float64 `json:"seek"`
			} `json:"videos"`
		} `json:"result"`
	} `json:"data"`
}

func (c *JumpProxyClient) GetSearchImg(req *GetSearchImgReq) (res GetSearchImgRes, err error) {
	err = c.request(req, &res)
	return
}
