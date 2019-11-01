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
	UserID   string `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	DataName string `form:"data_name" json:"data_name" xml:"data_name"  binding:"required"`
}

// RetListReq Binding from JSON
type RetListReq struct {
	UserID   string `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
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
		// json parse sect
		var json Test
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID, err := strconv.Atoi(json.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		password := json.Password
		dataName := json.DataName
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
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}

func handleError(err error) {
	if err !=  nil {
		panic("error")
	}
}
