package response

import (
	"github.com/gin-gonic/gin"
)

// Send ...
func Send(c *gin.Context, status int, resp interface{}) {
	accept := c.GetHeader("Accept")

	switch accept {
	case gin.MIMEXML, gin.MIMEXML2:
		c.XML(status, resp)
		break
	default:
		c.JSON(status, resp)
	}
}
