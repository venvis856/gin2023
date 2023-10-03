package response

import (
	"gin/internal/global/errcode"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 统一定义系统状态码
const (
	//http 返回结果
	HTTP_SUCCESS = 200
	HTTP_FAILURE = 500
	trackId      = "TrackId"
)

type ResponseJson struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
	TrackId string `json:"track_id"`
}

type PageResult struct {
	Total   int64 `json:"total"`
	Records any   `json:"records"`
}

// Json
func Json(c *gin.Context, httpCode, code int, msg string, data interface{}) {
	c.JSON(httpCode, ResponseJson{Code: code, Msg: msg, TrackId: TrackId(c), Data: data})
}

func Success(c *gin.Context, data any) {
	c.JSON(HTTP_SUCCESS, ResponseJson{Code: int(errcode.SUCCESS), Msg: errcode.SUCCESS.String(), TrackId: TrackId(c), Data: data})
}

func Error(c *gin.Context, err error) {
	code := int(errcode.ERROR)
	if v, ok := err.(errcode.ErrCode); ok {
		code = int(v)
	} else if v, ok := err.(errcode.ErrWrap); ok {
		code = v.ErrCode()
	}
	c.JSON(HTTP_SUCCESS, ResponseJson{Code: code, Msg: err.Error(), TrackId: TrackId(c), Data: nil})
}

func Data(c *gin.Context, err error, data any) {
	if err != nil {
		Error(c, err)
	} else {
		Success(c, data)
	}
}

func PageData(c *gin.Context, total int64, data any, err error) {
	if err != nil {
		Error(c, err)
	} else {
		Success(c, &PageResult{
			Total:   total,
			Records: data,
		})
	}
}
func ErrorWithData(c *gin.Context, err error, data interface{}) {
	code := int(errcode.ERROR)
	if v, ok := err.(errcode.ErrCode); ok {
		code = int(v)
	} else if v, ok := err.(errcode.ErrWrap); ok {
		code = v.ErrCode()
	}
	c.JSON(HTTP_SUCCESS, ResponseJson{Code: code, Msg: err.Error(), TrackId: TrackId(c), Data: data})
}

func TrackId(c *gin.Context) string {
	return c.GetString(trackId)
}

func SetTrackId(c *gin.Context) string {
	if _, ok := c.Keys[trackId]; ok {
		return c.GetString(trackId)
	}
	_uuid, _ := uuid.NewRandom()
	c.Set(trackId, _uuid.String())
	return _uuid.String()
}

func Fail(c *gin.Context, err error) {
	c.AbortWithError(HTTP_FAILURE, err)
}
