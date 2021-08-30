package router

import (
	. "webDemoBackend/internal/service"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func InitRouter() {
	r := gin.Default()

	routerLib := r.Group("/api/v1/library")
	routerLib.GET("/book", GetBook)
	routerLib.POST("/book", PostBook)
	routerLib.DELETE("/book", DeleteBook) //name ID
	routerLib.PATCH("/book", PatchBook)
	routerLib.PUT("/book", PutBook)

	r.Run()
}

// [GET] http://mytube.com/v1/videos/ -> 取得 video 列表
// [POST] http://mytube.com/v1/videos/ -> 新增 video
// [GET] http://mytube.com/v1/videos/MgphHyGgeQU -> 取得指定 ID 的 video
// [PUT] http://mytube.com/v1/videos/MgphHyGgeQU -> 修改指定 ID 的 video
// [DELETE] http://mytube.com/v1/videos/MgphHyGgeQU -> 刪除指定 ID 的 video
