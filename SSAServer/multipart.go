package ssa

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"strings"
	"os"

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
		case "password":
			st := string(slurp)
			ret.Password = &st
		default:
			mes := part.FormName() + " is invalid form data"
			return fmt.Errorf(mes)
		}
	}
	*p = server.NewSaveDataPayload(ret,"gr")
	return nil
}

// SSAServerSaveDataEncoderFunc implements the multipart encoder for service
// "SSAServer" endpoint "Save_data".
func SSAServerSaveDataEncoderFunc(mw *multipart.Writer, p *ssaserver.SaveDataPayload) error {
	// Add multipart request encoder logic here
	fmt.Println("called")
	return nil
}

// SSAServerPickUpDataDecoderFunc implements the multipart decoder for service
// "SSAServer" endpoint "Pick_up_data". The decoder must populate the argument
// p after encoding.
func SSAServerPickUpDataDecoderFunc(mr *multipart.Reader, p **ssaserver.PickUpDataPayload) error {
	// Add multipart request decoder logic here
	ret := &server.PickUpDataRequestBody{}
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
		case "image_name":
			st := string(slurp)
			ret.ImageName = &st
		default:
			mes := part.FormName() + " is invalid form data"
			return fmt.Errorf(mes)
		}
	}
	*p = server.NewPickUpDataPayload(ret,"hoge",1234)
	return nil
}

// SSAServerPickUpDataEncoderFunc implements the multipart encoder for service
// "SSAServer" endpoint "Pick_up_data".
func SSAServerPickUpDataEncoderFunc(mw *multipart.Writer, p *ssaserver.PickUpDataPayload) error {
	err := createFileField(mw, p, p.DataName, "Data")
	if err != nil {
		return err
	}
	err = mw.WriteField("data_type", strconv.Itoa(p.DataType))
	if err != nil {
		return err
	}
	err = mw.WriteField("data_name", p.DataName)
	if err != nil {
		return err
	}
	if p.DataType == 1 {
		err = mw.WriteField("title", "たいとる")
		if err != nil {
			return err
		}
		if !(p.ImageName == nil || strings.EqualFold(*p.ImageName, "")){
			err = createFileField(mw, p, *p.ImageName, "Image")
			if err != nil {
				return err
			}
			err = mw.WriteField("image_name", *p.ImageName)
			if err != nil {
				return err
			}
		}
	}
	err = mw.Close()
	if err != nil {
		return err
	}
	return nil
}

func createFileField(mw *multipart.Writer, p *ssaserver.PickUpDataPayload, name string, fname string) error {
	path := GetPickUpPath(p.GroupID, p.DataUserID, name)
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	file.Close()
	part, err := mw.CreateFormFile(fname, fi.Name())
	if err != nil {
		return err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return err
	}
	return nil
}
