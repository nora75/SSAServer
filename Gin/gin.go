package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "Gin/db"
	files "Gin/files"
	"os"
	"io/ioutil"
	"bufio"
	"strconv"
)


// Test Binding from JSON
type Test struct {
	Mail   string `form:"mail" json:"mail" xml:"mail"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// RetListReq Binding from JSON
type RetListReq struct {
	UserID   string `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/Login", func(c *gin.Context) {
		// json parse sect
		var json Test
		if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mail := json.Mail
		password := json.Password

		// ログイン情報の成否判定
		err := db.UserAuth(mail, password)
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := db.RetUserStruct(mail)
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_id": user.ID,
			"group_id": user.GroupID,
			"mail": mail,
			"password": password,
			"user_name": user.UserName,
		})
	})
	r.GET("/group/:group_id", func(c *gin.Context) {
		groupID := c.Param("group_id")
		userID, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		password := c.Query("password")
		err = db.GroupCheck(userID, password, groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		datalist, err := db.FindAllDataInGroup(groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, datalist)
	})
	r.GET("/group/:group_id/:data_user_id", func(c *gin.Context) {
		groupID := c.Param("group_id")
		dataUserID, err := strconv.Atoi(c.Param("data_user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		password := c.Query("password")
		dataName := c.Query("data_name")
		err = db.GroupCheck(userID, password, groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		path, err := files.GetPickUpPath(groupID,dataUserID,dataName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		file, err := os.Open(path)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fi, err := file.Stat()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()
		contentLength := fi.Size()
		contentType := http.DetectContentType(b)

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="`+dataName+`"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, bufio.NewReader(file), extraHeaders)
	})
	r.Run(":50113") // listen and serve on 0.0.0.0:50113

}

func handleError(err error) {
	if err !=  nil {
		panic("error")
	}
}
