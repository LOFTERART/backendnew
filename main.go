package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"upload/api"
	"upload/glo"
	"upload/middleware"
	"upload/model"
	"upload/util"
)

func init() {
	model.Initialized()
}

func init() {
	glo.IPAddress, _ = util.GetIPV4()
}

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.Use(middleware.Cors)
	//router.LoadHTMLFiles("./dist/index.html")
	//router.Static("/dist", "./dist")
	//router.GET("/admin", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "index.html", nil)
	//})
	//router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	//
	admin := router.Group("/adminService")
	{
		//获取file
		admin.POST("/get_file_list", api.GetFileList)
		//分享
		admin.POST("/share", api.ShareFile)
		//删除一个
		admin.POST("/del_one", api.DelFile)
	}

	Chunk := router.Group("/upload")
	{

		Chunk.GET("/checkChunk", api.CheckChunk)

		Chunk.POST("/uploadChunk", api.UploadChunk)

		Chunk.GET("/meagerChunk", api.MeagerChunk)
	}

	router.Use(middleware.FileFilterMiddle)
	router.StaticFS("/uploadFile", http.Dir("./uploadFile"))
	_ = router.Run(":8080")
}
