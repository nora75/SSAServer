package ssa

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	ssaserver "SSAServer/gen/ssa_server"
	// client "SSAServer/gen/http/ssa_server/client"
	server "SSAServer/gen/http/ssa_server/server"
)

// SSAServerSaveDataDecoderFunc implements the multipart decoder for service
// "SSAServer" endpoint "Save_data". The decoder must populate the argument p
// after encoding.
func SSAServerSaveDataDecoderFunc(mr *multipart.Reader, p **ssaserver.SaveDataPayload) error {
	var ret *server.SaveDataRequestBody
	// Add multipart request decoder logic here
	fmt.Println("bef for")
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to load part: %s", err)
		}
		fmt.Println("after err nil")
		if part.FormName() == "user_id" {
			fmt.Println(part.Header.Get("user_id"))
			// ret.UserID = part.Header.Get()
		} else if part.FormName() == "data_name" {
			fmt.Println(part.Header.Get("data_name"))
			*ret.DataName = part.Header.Get("image_name")
		} else if part.FormName() == "data_type" {
			fmt.Println(part.Header.Get("data_type"))
			*ret.DataType,_ = strconv.Atoi(part.Header.Get("data_type"))
		} else if part.FormName() == "title" {
			fmt.Println(part.Header.Get("title"))
			*ret.Title = part.Header.Get("title")
		} else if part.FormName() == "image_name" {
			fmt.Println(part.Header.Get("image_name"))
			*ret.ImageName = part.Header.Get("image_name")
		} else if part.FormName() == "Data" {
			dataFile, err := os.Create(*ret.DataName)
			if err != nil {
				return fmt.Errorf("failed to open file: %s", err)
			}
			io.Copy(dataFile, part)
		} else if part.FormName() == "Image" {
			imageFile, err := os.Create(*ret.ImageName)
			if err != nil {
				return fmt.Errorf("failed to open file: %s", err)
			}
			io.Copy(imageFile, part)
		} else {
			mes := part.FormName() + "' is invalid form data"
			return fmt.Errorf(mes)
		}
	}
	*p = server.NewSaveDataPayload(ret,"hoge")
	return nil
}

// SSAServerSaveDataEncoderFunc implements the multipart encoder for service
// "SSAServer" endpoint "Save_data".
func SSAServerSaveDataEncoderFunc(mw *multipart.Writer, p *ssaserver.SaveDataPayload) error {
	// Add multipart request encoder logic here
	return nil
}

// SSAServerPickUpDataDecoderFunc implements the multipart decoder for service
// "SSAServer" endpoint "Pick_up_data". The decoder must populate the argument
// p after encoding.
func SSAServerPickUpDataDecoderFunc(mr *multipart.Reader, p **ssaserver.PickUpDataPayload) error {
	// Add multipart request decoder logic here
	return nil
}

// SSAServerPickUpDataEncoderFunc implements the multipart encoder for service
// "SSAServer" endpoint "Pick_up_data".
func SSAServerPickUpDataEncoderFunc(mw *multipart.Writer, p *ssaserver.PickUpDataPayload) error {
	// Add multipart request encoder logic here
	return nil
}

/*
/ Make new file by filename var
*/
// func MakeFile(fileName string) *os.File {
// 	fileVar, err := os.Create(fileName)
// 	if err != nil {
// 		return fmt.Errorf("failed to open file: %s", err)
// 	}
// 	return fileVar
// }
