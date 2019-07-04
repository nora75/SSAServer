package ssa

import (
	"mime/multipart"
	"io"
	"os"
	"fmt"

	// client "SSAServer/gen/http/ssa_server/client"
	// server "SSAServer/gen/http/ssa_server/server"
	ssaserver "SSAServer/gen/ssa_server"
)

// SSAServerSaveDataDecoderFunc implements the multipart decoder for service
// "SSAServer" endpoint "Save_data". The decoder must populate the argument p
// after encoding.
func SSAServerSaveDataDecoderFunc(mr *multipart.Reader, p **ssaserver.SaveDataPayload) error {
	// Add multipart request decoder logic here
	file, err := os.Create("testhoge")
	if err != nil {
		return fmt.Errorf("failed to load part: %s", err)
	}
	fmt.Println("before loop")
	for {
		fmt.Println("start")
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to load part: %s", err)
		}
		io.Copy(file, part)
	}
	fmt.Println("after loop")
	defer file.Close()
	fmt.Println("after loop")
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
