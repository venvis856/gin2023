package common_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		_uuid, _ := uuid.NewRandom()
		uuidStr := _uuid.String()
		c.Header("trace_id", uuidStr)
		c.Set("trace_id", uuidStr)
		c.Next()
	}
}
