package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"io/ioutil"
	"io"
	"bufio"
	// "bytes"
	"fmt"
	"mime/multipart"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/group/:group_id/:data_user_id", func(c *gin.Context) {
		groupID := c.Param("group_id")
		dataUserID := c.Param("data_user_id")
		dataName := c.PostForm("data_name")
		// dataType := c.PostForm("data_type")
		userID := c.PostForm("user_id")
		// imageName := c.DefaultPostForm("image_name", "")
		path := GetPickUpPath(groupID,dataUserID,dataName)
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(os.Stderr, err)
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

		head := `attachment; filename="`+dataName+`"`
		extraHeaders := map[string]string{
			"Content-Disposition": head,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, bufio.NewReader(file), extraHeaders)
	})
	r.GET("/PickUp", func(c *gin.Context) {
		// path := GetPickUpPath(groupID, UserID, dataName,)
		path := "/home/nora75/go/src/SSAServer/group-h/1/ansible.md"
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(os.Stderr, err)
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
			"Content-Disposition": `attachment; filename="ansible.md"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, bufio.NewReader(file), extraHeaders)
	})

	r.GET("/somedata", func(c *gin.Context) {
		pr, pw := io.Pipe()
		// io.PipeWriterをmultipart.Writerに渡す
		mw := multipart.NewWriter(pw)

		// Goルーチンを開始し非同期でpwに書き込む
		go func() {
			// prの処理が正常に終わるように必ずpwをクローズする
			defer pw.Close()
			defer mw.Close()
			err := mw.WriteField("data_type", "1")
			handleError(err)
			err = mw.WriteField("title", "たいとる")
			handleError(err)
			err = mw.WriteField("data_name", "ansible.md")
			handleError(err)
			filename := "/home/nora75/go/src/SSAServer/group-h/1/ansible.md"
			err = hoge(mw, filename, "data")
			handleError(err)
			err = mw.WriteField("image_name", "hoge.md")
			handleError(err)
			filename = "/home/nora75/go/src/SSAServer/group-h/1/hoge.md"
			err = hoge(mw, filename, "image")
			handleError(err)

		}()

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="ansible.md"; filename="hoge.md"`,
		}

		// buf := &bytes.Buffer{}
		// nRead, err := io.Copy(buf, *pr) 
		// nioutil.ReadAll(*pr)
		// if err != nil {
		// 	    fmt.Println(err)
				
		// }
		// contentLength := nRead

		c.DataFromReader(http.StatusOK, 0, mw.FormDataContentType(), pr, extraHeaders)

	})

	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return

		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,

		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}

func hoge(mw *multipart.Writer, path string, field string) (err error) {
	file, err := os.Open(path)
	handleError(err)
	fi, err := file.Stat()
	handleError(err)
	fw, err := mw.CreateFormFile(field, fi.Name())
	handleError(err)
	if _, err := io.Copy(fw, file); err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func handleError(err error) {
	if err !=  nil {
		panic("error")
	}
}
