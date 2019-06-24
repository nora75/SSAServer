package ssa

import (
	ssaserver "SSAServer/gen/ssa_server"
	"mime/multipart"
)

// SSAServerSaveDataDecoderFunc implements the multipart decoder for service
// "SSAServer" endpoint "Save_data". The decoder must populate the argument p
// after encoding.
func SSAServerSaveDataDecoderFunc(mr *multipart.Reader, p **ssaserver.SaveDataPayload) error {
	// Add multipart request decoder logic here
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
