package ssa

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	// "strings"

	ssaserver "SSAServer/gen/ssa_server"
	// client "SSAServer/gen/http/ssa_server/client"
	server "SSAServer/gen/http/ssa_server/server"
)

// SSAServerSaveDataDecoderFunc implements the multipart decoder for service
// "SSAServer" endpoint "Save_data". The decoder must populate the argument p
// after encoding.
func SSAServerSaveDataDecoderFunc(mr *multipart.Reader, p **ssaserver.SaveDataPayload) error {
	ret := &server.SaveDataRequestBody{}
	// Add multipart request decoder logic here
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to load part: %s", err)
		}
		slurp, err := ioutil.ReadAll(part)
		if err != nil {
			return fmt.Errorf("failed to load part: %s", err)
		}
		switch part.FormName() {
		case "user_id":
			st, _ := strconv.Atoi(string(slurp))
			ret.UserID = &st
		case "data_name":
			st := string(slurp)
			ret.DataName = &st
		case "data_type":
			st, _ := strconv.Atoi(string(slurp))
			ret.DataType = &st
		case "title":
			st := string(slurp)
			ret.Title = &st
		case "image_name":
			st := string(slurp)
			ret.ImageName = &st
		case "Data":
			ret.Data = slurp
		case "Image":
			ret.Image = slurp
		default:
			break
		}
	}
	*p = server.NewSaveDataPayload(ret,"hoge")
	return nil
}

// SSAServerSaveDataEncoderFunc implements the multipart encoder for service
// "SSAServer" endpoint "Save_data".
func SSAServerSaveDataEncoderFunc(mw *multipart.Writer, p *ssaserver.SaveDataPayload) error {
	// Add multipart request encoder logic here
	// dataFile, err := os.Create(p.DataName)
	// defer func() {
	// 	if err := dataFile.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// if err != nil {
	// 	return fmt.Errorf("failed to open file: %s", err)
	// }
	// _, err = io.Copy(dataFile, *p.Data)
	// if err != nil {
	// 	return fmt.Errorf("failed to open file: %s", err)
	// }
	// if p.Image == nil || strings.EqualFold(*p.ImageName, ""){
	// 	imageFile, err := os.Create(*p.ImageName)
	// 	defer func() {
	// 		if err := imageFile.Close(); err != nil {
	// 			panic(err)
	// 		}
	// 	}()
	// 	if err != nil {
	// 		return fmt.Errorf("failed to open file: %s", err)
	// 	}
	// 	_, err = io.Copy(imageFile, *p.Image)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to open file: %s", err)
	// 	}
	// }
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
