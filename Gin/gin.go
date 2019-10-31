package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "Gin/db"
	files "Gin/files"
	"os"
	"io/ioutil"
	"bufio"
	"fmt"
	"strconv"
)


// Test Binding from JSON
type Test struct {
	UserID   string `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	DataName string `form:"data_name" json:"data_name" xml:"data_name"  binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/group/:group_id/:data_user_id", func(c *gin.Context) {
		groupID := c.Param("group_id")
		dataUserID, err := strconv.Atoi(c.Param("data_user_id"))
		if err != nil {
			fmt.Println("error in datauser id")
			fmt.Println(err)
			os.Exit(1)
		}
		// json parse sect
		var json Test
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID, err := strconv.Atoi(json.UserID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		password := json.Password
		dataName := json.DataName
		err = db.GroupCheck(userID, password, groupID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path, err := files.GetPickUpPath(groupID,dataUserID,dataName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file, err := os.Open(path)
		if err != nil {
			panic("hoge")
		}
		fi, err := file.Stat()
		if err != nil {
			panic("hoge")
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
