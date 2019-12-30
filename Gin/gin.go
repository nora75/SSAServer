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
	"time"
	"math/rand"
	"strings"
)

// for RandString
var randSrc = rand.NewSource(time.Now().UnixNano())

// for RandString
const (
	rs6Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rs6LetterIdxBits = 6
	rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
	rs6LetterIdxMax  = 63 / rs6LetterIdxBits
)

// LoginJSON Binding from JSON
type LoginJSON struct {
	Mail   string `form:"mail" json:"mail" xml:"mail"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// ChangeJSON Binding from JSON
type ChangeJSON struct {
	GroupID   string `form:"group_id" json:"group_id" xml:"group_id" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// DeleteJSON Binding from JSON
type DeleteJSON struct {
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// RegistrationJSON Binding from JSON
type RegistrationJSON struct {
	GroupID   string `form:"group_id" json:"group_id" xml:"group_id"`
	Mail   string `form:"mail" json:"mail" xml:"mail"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	UserName   string `form:"user_name" json:"user_name" xml:"user_name"  binding:"required"`
}

// RetListJSON Binding from JSON
type RetListJSON struct {
	UserID   string `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/Registration", func(c *gin.Context) {
		var json RegistrationJSON
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var groupID string
		if (strings.EqualFold(json.GroupID, "")) {
			groupID = "group-" + RandString(12)
		} else {
			groupID = json.GroupID
		}
		mail := json.Mail
		password := json.Password
		userName := json.UserName

		userID, err := db.InsertUserData(userName, password, mail, groupID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = files.CreateUserDir(groupID, userID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
			"group_id": groupID,
			"mail": mail,
			"user_name": userName,

		})
	})
	r.POST("/Login", func(c *gin.Context) {
		// json parse sect
		var json LoginJSON
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
			"user_name": user.UserName,
			// "mail": mail,
			// "password": password,

		})
	})
	r.POST("/users/:user_id", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var json ChangeJSON
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		groupID := json.GroupID
		password := json.Password

		var oldGroupID string
		oldGroupID, err = db.UpdateGroupID(userID, groupID, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = files.MoveUserDir(oldGroupID, groupID, userID)
		if err != nil {
			_, err = db.UpdateGroupID(userID, oldGroupID, password)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": true,
		})
	})
	r.DELETE("/users/:user_id", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var json DeleteJSON
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		password := json.Password
		err = db.DeleteUser(userID, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": true,
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
	r.POST("/group/:group_id", func(c *gin.Context) {
		groupID := c.Param("group_id")
		dataType, err := strconv.Atoi(c.PostForm("data_type"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dataName := c.PostForm("data_name")
		imageName := c.PostForm("image_name")
		title := c.PostForm("title")
		var userID int
		userID, err = strconv.Atoi(c.PostForm("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		password := c.PostForm("password")
		switch dataType {
		case 0:
		case 1:
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		file, err := c.FormFile("Data")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		filedata, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := ioutil.ReadAll(filedata)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer filedata.Close()

		var image []byte
		if (dataType == 1 && file != nil) {
			file, _ = c.FormFile("Image")
			filedata, err := file.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			image, err = ioutil.ReadAll(filedata)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			defer filedata.Close()

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

				return
			}

		}
		err = db.InsertDataData(userID, password, groupID, dataName, imageName, title, dataType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = files.SaveFile(data, groupID, userID, dataName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if (dataType == 1) {
			if !(image == nil || strings.EqualFold(imageName, "")) {
				err = files.SaveFile(image, groupID, userID, imageName)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

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


// RandString Return Random n length String
func RandString(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rs6LetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rs6LetterIdxMax
		}
		idx := int(cache & rs6LetterIdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i--
		}
		cache >>= rs6LetterIdxBits
		remain--
	}
	return string(b)
}
