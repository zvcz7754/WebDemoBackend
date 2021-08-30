package service

import (
	"fmt"
	"net/http"
	. "webDemoBackend/internal/entity"
	. "webDemoBackend/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	// c.JSON(200, Book{})

	id := c.Query("id")
	fmt.Println("get id =>", id)

}

func PostBook(c *gin.Context) { //取比對若已經有則返回409
	fmt.Println("post body =>")

}

func DeleteBook(c *gin.Context) {
	id := c.Params["id"]
	fmt.Println("delete id =>", id)
}

func PatchBook(c *gin.Context) {
}

func PutBook(c *gin.Context) {
}
