package controller

import (
	"go-hello/go/entity"
	"go-hello/go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/GetUserList [get]
// @Param id path int false "使用者 id"
func GetUserById(c *gin.Context) {
	todoList, err := service.GetUserById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todoList,
		})
	}
}

// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/GetUserList [get]
func GetUserList(c *gin.Context) {
	todoList, err := service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todoList,
		})
	}
}

// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/CreateUser [post]
// @Param name formData string false "使用者姓名"
func CreateUser(c *gin.Context) {
	//定義一個User變數
	var user entity.User
	//將呼叫後端的request請求中的body資料根據json格式解析到User結構變數中
	c.BindJSON(&user)
	//將被轉換的user變數傳給service層的CreateUser方法，進行User的新建
	err := service.CreateUser(&user)
	//判斷是否異常，無異常則返回包含200和更新資料的資訊
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}

// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/GetUserList [get]
// @Param id query int false "使用者 id"
func UpdateUser(c *gin.Context) {

}

// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/GetUserList [get]
// @Param id path int false "使用者 id"
func DeleteUserById(c *gin.Context) {

}

// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/name [get]
// @Param name query string false "使用者姓名"
// func getName(c *gin.Context) {
// 	if name := c.Query("name"); name == "" {
// 		c.Data(http.StatusBadRequest, "text/plain; charset=utf-8;", []byte("Please enter your name!"))
// 	} else {
// 		client.Connect()
// 		// _, err := client.Get()
// 		// if err == nil {
// 		// 	c.Data(http.StatusOK, "text/plain; charset=utf-8;", []byte("hello"+name))
// 		// }
// 		// c.Data(http.StatusOK, "text/plain; charset=utf-8;", []byte("hello"+client.Get()))
// 		client.Disconnect()
// 		// name := c.DefaultQuery("name", "")
// 	}
// }

// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/name2 [get]
// @Param name path string false "使用者姓名"
// func getName2(c *gin.Context) {
// 	if name := c.Query("name"); name == "" {
// 		c.String(http.StatusBadRequest, "text/plain; charset=utf-8;", "Please enter your name!")
// 	} else {
// 		c.String(http.StatusOK, "text/plain; charset=utf-8;", "hello"+name)
// 		client.Connect()

// 		client.Disconnect()
// 		// name := c.DefaultQuery("name", "")
// 	}
// }

// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/name2 [put]
// @Param name path string false "使用者姓名"
// func updateName(c *gin.Context) {
// 	client.Connect()

// 	client.Disconnect()
// }

// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/name2 [delete]
// @Param name query string false "使用者姓名"
// func deleteName(c *gin.Context) {
// 	client.Connect()

// 	client.Disconnect()
// }

// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /demo/v1/postName [post]
// @Param name formData string false "使用者姓名"
// func postName(c *gin.Context) {
// 	var user entity.User
// 	c.BindJSON(&user)
// 	err := service.Insert(&user)
// 	name := c.PostForm("name")
// 	// name := c.DefaultPostForm("name", "Joe") // DefaultPostForm 取不到值時會返回指定的默認值
// 	if name == "" {
// 		c.JSON(500, gin.H{
// 			"status": gin.H{
// 				"code":   500,
// 				"status": "Internal Server Error",
// 			},
// 			"msg": "retry:true",
// 		})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": gin.H{
// 				"code":   http.StatusOK,
// 				"status": "成功",
// 			},
// 			"msg": "hello " + name,
// 		})
// 	}
// 	client.Connect()
// user := dao.User{
// 	Name: name,
// }
// client.Insert(user)
// players, err := client.Get()
// if err != nil {
//  fmt.Println(err)
// }
// if len(players) > 1 {
// 	players[1].Budget = 2000
// 	client.Update(players[1])
//  }

// users, err = client.Get()
// c.JSON(http.StatusOK, gin.H{
// 	"user": gin.H{
// 		"id":   user.Id,
// 		"name": user.Name,
// 	},
// })
// client.Disconnect()
