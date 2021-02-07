/**
 * @Author: LOFTER
 * @Description:
 * @File:  middle
 * @Date: 2021/2/5 3:04 下午
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"upload/util"
)

func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	c.Next()
}

//文件过滤

func FileFilterMiddle(c *gin.Context) {
	// 校验token 和 过期时间
	token := c.Request.URL.Query().Get("token")

	b, _ := util.DePwdCode(token)

	time1, _ := strconv.Atoi(string(b))

	if time.Now().UnixNano() <= int64(time1) {
		c.Next()
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "禁止下载",
		})
		c.Abort()
	}

}
