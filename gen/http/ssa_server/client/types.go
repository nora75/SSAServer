// Code generated by goa v3.0.3, DO NOT EDIT.
//
// SSAServer HTTP client types
//
// Command:
// $ goa gen SSAServer/design

package client

import (
	ssaserver "SSAServer/gen/ssa_server"
	ssaserverviews "SSAServer/gen/ssa_server/views"

	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "SSAServer" service "Register"
// endpoint HTTP request body.
type RegisterRequestBody struct {
	// User Name
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	// User Password
	Password string `form:"password" json:"password" xml:"password"`
	// User mail-address
	Mail string `form:"mail" json:"mail" xml:"mail"`
	// Group ID
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" xml:"group_id,omitempty"`
}

// LoginRequestBody is the type of the "SSAServer" service "Login" endpoint
// HTTP request body.
type LoginRequestBody struct {
	// User mail-address
	Mail string `form:"mail" json:"mail" xml:"mail"`
	// User Password
	Password string `form:"password" json:"password" xml:"password"`
}

// ChangeGroupRequestBody is the type of the "SSAServer" service "Change_group"
// endpoint HTTP request body.
type ChangeGroupRequestBody struct {
	// User password
	Password string `form:"password" json:"password" xml:"password"`
	// Group ID
	GroupID string `form:"group_id" json:"group_id" xml:"group_id"`
}

// DeleteUserRequestBody is the type of the "SSAServer" service "Delete_user"
// endpoint HTTP request body.
type DeleteUserRequestBody struct {
	// User Password
	Password string `form:"password" json:"password" xml:"password"`
}

// SaveDataRequestBody is the type of the "SSAServer" service "Save_data"
// endpoint HTTP request body.
type SaveDataRequestBody struct {
	// User ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
	// Data name
	DataName string `form:"data_name" json:"data_name" xml:"data_name"`
	// Data name
	DataType int `form:"data_type" json:"data_type" xml:"data_type"`
	// Data
	Data []byte `form:"Data" json:"Data" xml:"Data"`
	// Diary title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Image name
	ImageName *string `form:"image_name,omitempty" json:"image_name,omitempty" xml:"image_name,omitempty"`
	// Image
	Image []byte `form:"Image,omitempty" json:"Image,omitempty" xml:"Image,omitempty"`
}

// ReturnDataListRequestBody is the type of the "SSAServer" service
// "Return_data_list" endpoint HTTP request body.
type ReturnDataListRequestBody struct {
	// User ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// PickUpDataRequestBody is the type of the "SSAServer" service "Pick_up_data"
// endpoint HTTP request body.
type PickUpDataRequestBody struct {
	// User ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
	// Data name
	DataName string `form:"data_name" json:"data_name" xml:"data_name"`
}

// RegisterResponseBody is the type of the "SSAServer" service "Register"
// endpoint HTTP response body.
type RegisterResponseBody struct {
	// User id
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User Name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// User Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// User mail-address
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// Group ID
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" xml:"group_id,omitempty"`
	// Data name
	DataName *string `form:"data_name,omitempty" json:"data_name,omitempty" xml:"data_name,omitempty"`
	// Data
	Data []byte `form:"Data,omitempty" json:"Data,omitempty" xml:"Data,omitempty"`
	// Image data
	Image []byte `form:"Image,omitempty" json:"Image,omitempty" xml:"Image,omitempty"`
	// Data's name
	DataType *int `form:"data_type,omitempty" json:"data_type,omitempty" xml:"data_type,omitempty"`
	// Data title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// image's name
	ImageName *string `form:"image_name,omitempty" json:"image_name,omitempty" xml:"image_name,omitempty"`
	// date time
	DateTime *string `form:"date_time,omitempty" json:"date_time,omitempty" xml:"date_time,omitempty"`
}

// ReturnDataListResponseBody is the type of the "SSAServer" service
// "Return_data_list" endpoint HTTP response body.
type ReturnDataListResponseBody []*SsaResultResponse

// PickUpDataResponseBody is the type of the "SSAServer" service "Pick_up_data"
// endpoint HTTP response body.
type PickUpDataResponseBody struct {
	// User id
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User Name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// User Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// User mail-address
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// Group ID
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" xml:"group_id,omitempty"`
	// Data name
	DataName *string `form:"data_name,omitempty" json:"data_name,omitempty" xml:"data_name,omitempty"`
	// Data
	Data []byte `form:"Data,omitempty" json:"Data,omitempty" xml:"Data,omitempty"`
	// Image data
	Image []byte `form:"Image,omitempty" json:"Image,omitempty" xml:"Image,omitempty"`
	// Data's name
	DataType *int `form:"data_type,omitempty" json:"data_type,omitempty" xml:"data_type,omitempty"`
	// Data title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// image's name
	ImageName *string `form:"image_name,omitempty" json:"image_name,omitempty" xml:"image_name,omitempty"`
	// date time
	DateTime *string `form:"date_time,omitempty" json:"date_time,omitempty" xml:"date_time,omitempty"`
}

// RegisterInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Register" endpoint HTTP response body for the "Invalid_Group_ID" error.
type RegisterInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RegisterTheUserAlreadyExistsResponseBody is the type of the "SSAServer"
// service "Register" endpoint HTTP response body for the
// "The_user_already_exists" error.
type RegisterTheUserAlreadyExistsResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RegisterInvalidRequestResponseBody is the type of the "SSAServer" service
// "Register" endpoint HTTP response body for the "Invalid_Request" error.
type RegisterInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LoginInvalidRequestResponseBody is the type of the "SSAServer" service
// "Login" endpoint HTTP response body for the "Invalid_Request" error.
type LoginInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LoginInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Login" endpoint HTTP response body for the "Invalid_Group_ID" error.
type LoginInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ChangeGroupInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Change_group" endpoint HTTP response body for the "Invalid_Group_ID" error.
type ChangeGroupInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ChangeGroupInvalidRequestResponseBody is the type of the "SSAServer" service
// "Change_group" endpoint HTTP response body for the "Invalid_Request" error.
type ChangeGroupInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteUserInvalidRequestResponseBody is the type of the "SSAServer" service
// "Delete_user" endpoint HTTP response body for the "Invalid_Request" error.
type DeleteUserInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SaveDataInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Save_data" endpoint HTTP response body for the "Invalid_Group_ID" error.
type SaveDataInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SaveDataInvalidRequestResponseBody is the type of the "SSAServer" service
// "Save_data" endpoint HTTP response body for the "Invalid_Request" error.
type SaveDataInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SaveDataInvalidDataResponseBody is the type of the "SSAServer" service
// "Save_data" endpoint HTTP response body for the "Invalid_Data" error.
type SaveDataInvalidDataResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReturnDataListInvalidGroupIDResponseBody is the type of the "SSAServer"
// service "Return_data_list" endpoint HTTP response body for the
// "Invalid_Group_ID" error.
type ReturnDataListInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReturnDataListInvalidRequestResponseBody is the type of the "SSAServer"
// service "Return_data_list" endpoint HTTP response body for the
// "Invalid_Request" error.
type ReturnDataListInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PickUpDataInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Pick_up_data" endpoint HTTP response body for the "Invalid_Group_ID" error.
type PickUpDataInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PickUpDataInvalidRequestResponseBody is the type of the "SSAServer" service
// "Pick_up_data" endpoint HTTP response body for the "Invalid_Request" error.
type PickUpDataInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PickUpDataInvalidDataNameResponseBody is the type of the "SSAServer" service
// "Pick_up_data" endpoint HTTP response body for the "Invalid_data_name" error.
type PickUpDataInvalidDataNameResponseBody struct {
	// Return false when Error occur
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
	// Error returned.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SsaResultResponse is used to define fields on response body types.
type SsaResultResponse struct {
	// User id
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User Name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// User Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// User mail-address
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// Group ID
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" xml:"group_id,omitempty"`
	// Data name
	DataName *string `form:"data_name,omitempty" json:"data_name,omitempty" xml:"data_name,omitempty"`
	// Data
	Data []byte `form:"Data,omitempty" json:"Data,omitempty" xml:"Data,omitempty"`
	// Image data
	Image []byte `form:"Image,omitempty" json:"Image,omitempty" xml:"Image,omitempty"`
	// Data's name
	DataType *int `form:"data_type,omitempty" json:"data_type,omitempty" xml:"data_type,omitempty"`
	// Data title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// image's name
	ImageName *string `form:"image_name,omitempty" json:"image_name,omitempty" xml:"image_name,omitempty"`
	// date time
	DateTime *string `form:"date_time,omitempty" json:"date_time,omitempty" xml:"date_time,omitempty"`
}

// NewRegisterRequestBody builds the HTTP request body from the payload of the
// "Register" endpoint of the "SSAServer" service.
func NewRegisterRequestBody(p *ssaserver.RegisterPayload) *RegisterRequestBody {
	body := &RegisterRequestBody{
		UserName: p.UserName,
		Password: p.Password,
		Mail:     p.Mail,
		GroupID:  p.GroupID,
	}
	return body
}

// NewLoginRequestBody builds the HTTP request body from the payload of the
// "Login" endpoint of the "SSAServer" service.
func NewLoginRequestBody(p *ssaserver.LoginPayload) *LoginRequestBody {
	body := &LoginRequestBody{
		Mail:     p.Mail,
		Password: p.Password,
	}
	return body
}

// NewChangeGroupRequestBody builds the HTTP request body from the payload of
// the "Change_group" endpoint of the "SSAServer" service.
func NewChangeGroupRequestBody(p *ssaserver.ChangeGroupPayload) *ChangeGroupRequestBody {
	body := &ChangeGroupRequestBody{
		Password: p.Password,
		GroupID:  p.GroupID,
	}
	return body
}

// NewDeleteUserRequestBody builds the HTTP request body from the payload of
// the "Delete_user" endpoint of the "SSAServer" service.
func NewDeleteUserRequestBody(p *ssaserver.DeleteUserPayload) *DeleteUserRequestBody {
	body := &DeleteUserRequestBody{
		Password: p.Password,
	}
	return body
}

// NewSaveDataRequestBody builds the HTTP request body from the payload of the
// "Save_data" endpoint of the "SSAServer" service.
func NewSaveDataRequestBody(p *ssaserver.SaveDataPayload) *SaveDataRequestBody {
	body := &SaveDataRequestBody{
		UserID:    p.UserID,
		DataName:  p.DataName,
		DataType:  p.DataType,
		Data:      p.Data,
		Title:     p.Title,
		ImageName: p.ImageName,
		Image:     p.Image,
	}
	return body
}

// NewReturnDataListRequestBody builds the HTTP request body from the payload
// of the "Return_data_list" endpoint of the "SSAServer" service.
func NewReturnDataListRequestBody(p *ssaserver.ReturnDataListPayload) *ReturnDataListRequestBody {
	body := &ReturnDataListRequestBody{
		UserID: p.UserID,
	}
	return body
}

// NewPickUpDataRequestBody builds the HTTP request body from the payload of
// the "Pick_up_data" endpoint of the "SSAServer" service.
func NewPickUpDataRequestBody(p *ssaserver.PickUpDataPayload) *PickUpDataRequestBody {
	body := &PickUpDataRequestBody{
		UserID:   p.UserID,
		DataName: p.DataName,
	}
	return body
}

// NewRegisterSsaResultOK builds a "SSAServer" service "Register" endpoint
// result from a HTTP "OK" response.
func NewRegisterSsaResultOK(body *RegisterResponseBody) *ssaserverviews.SsaResultView {
	v := &ssaserverviews.SsaResultView{
		UserID:    body.UserID,
		UserName:  body.UserName,
		Password:  body.Password,
		Mail:      body.Mail,
		GroupID:   body.GroupID,
		DataName:  body.DataName,
		Data:      body.Data,
		Image:     body.Image,
		DataType:  body.DataType,
		Title:     body.Title,
		ImageName: body.ImageName,
		DateTime:  body.DateTime,
	}
	return v
}

// NewRegisterInvalidGroupID builds a SSAServer service Register endpoint
// Invalid_Group_ID error.
func NewRegisterInvalidGroupID(body *RegisterInvalidGroupIDResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewRegisterTheUserAlreadyExists builds a SSAServer service Register endpoint
// The_user_already_exists error.
func NewRegisterTheUserAlreadyExists(body *RegisterTheUserAlreadyExistsResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewRegisterInvalidRequest builds a SSAServer service Register endpoint
// Invalid_Request error.
func NewRegisterInvalidRequest(body *RegisterInvalidRequestResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewLoginInvalidRequest builds a SSAServer service Login endpoint
// Invalid_Request error.
func NewLoginInvalidRequest(body *LoginInvalidRequestResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewLoginInvalidGroupID builds a SSAServer service Login endpoint
// Invalid_Group_ID error.
func NewLoginInvalidGroupID(body *LoginInvalidGroupIDResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewChangeGroupInvalidGroupID builds a SSAServer service Change_group
// endpoint Invalid_Group_ID error.
func NewChangeGroupInvalidGroupID(body *ChangeGroupInvalidGroupIDResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewChangeGroupInvalidRequest builds a SSAServer service Change_group
// endpoint Invalid_Request error.
func NewChangeGroupInvalidRequest(body *ChangeGroupInvalidRequestResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewDeleteUserInvalidRequest builds a SSAServer service Delete_user endpoint
// Invalid_Request error.
func NewDeleteUserInvalidRequest(body *DeleteUserInvalidRequestResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewSaveDataInvalidGroupID builds a SSAServer service Save_data endpoint
// Invalid_Group_ID error.
func NewSaveDataInvalidGroupID(body *SaveDataInvalidGroupIDResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewSaveDataInvalidRequest builds a SSAServer service Save_data endpoint
// Invalid_Request error.
func NewSaveDataInvalidRequest(body *SaveDataInvalidRequestResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewSaveDataInvalidData builds a SSAServer service Save_data endpoint
// Invalid_Data error.
func NewSaveDataInvalidData(body *SaveDataInvalidDataResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewReturnDataListSsaResultCollectionOK builds a "SSAServer" service
// "Return_data_list" endpoint result from a HTTP "OK" response.
func NewReturnDataListSsaResultCollectionOK(body ReturnDataListResponseBody) ssaserverviews.SsaResultCollectionView {
	v := make([]*ssaserverviews.SsaResultView, len(body))
	for i, val := range body {
		v[i] = &ssaserverviews.SsaResultView{
			UserID:    val.UserID,
			UserName:  val.UserName,
			Password:  val.Password,
			Mail:      val.Mail,
			GroupID:   val.GroupID,
			DataName:  val.DataName,
			Data:      val.Data,
			Image:     val.Image,
			DataType:  val.DataType,
			Title:     val.Title,
			ImageName: val.ImageName,
			DateTime:  val.DateTime,
		}
	}
	return v
}

// NewReturnDataListInvalidGroupID builds a SSAServer service Return_data_list
// endpoint Invalid_Group_ID error.
func NewReturnDataListInvalidGroupID(body *ReturnDataListInvalidGroupIDResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewReturnDataListInvalidRequest builds a SSAServer service Return_data_list
// endpoint Invalid_Request error.
func NewReturnDataListInvalidRequest(body *ReturnDataListInvalidRequestResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewPickUpDataSsaResultOK builds a "SSAServer" service "Pick_up_data"
// endpoint result from a HTTP "OK" response.
func NewPickUpDataSsaResultOK(body *PickUpDataResponseBody) *ssaserverviews.SsaResultView {
	v := &ssaserverviews.SsaResultView{
		UserID:    body.UserID,
		UserName:  body.UserName,
		Password:  body.Password,
		Mail:      body.Mail,
		GroupID:   body.GroupID,
		DataName:  body.DataName,
		Data:      body.Data,
		Image:     body.Image,
		DataType:  body.DataType,
		Title:     body.Title,
		ImageName: body.ImageName,
		DateTime:  body.DateTime,
	}
	return v
}

// NewPickUpDataInvalidGroupID builds a SSAServer service Pick_up_data endpoint
// Invalid_Group_ID error.
func NewPickUpDataInvalidGroupID(body *PickUpDataInvalidGroupIDResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewPickUpDataInvalidRequest builds a SSAServer service Pick_up_data endpoint
// Invalid_Request error.
func NewPickUpDataInvalidRequest(body *PickUpDataInvalidRequestResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// NewPickUpDataInvalidDataName builds a SSAServer service Pick_up_data
// endpoint Invalid_data_name error.
func NewPickUpDataInvalidDataName(body *PickUpDataInvalidDataNameResponseBody) *ssaserver.SsaError {
	v := &ssaserver.SsaError{
		Fault:   *body.Fault,
		Message: *body.Message,
	}
	return v
}

// ValidateRegisterInvalidGroupIDResponseBody runs the validations defined on
// Register_Invalid_Group_ID_Response_Body
func ValidateRegisterInvalidGroupIDResponseBody(body *RegisterInvalidGroupIDResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRegisterTheUserAlreadyExistsResponseBody runs the validations
// defined on Register_The_user_already_exists_Response_Body
func ValidateRegisterTheUserAlreadyExistsResponseBody(body *RegisterTheUserAlreadyExistsResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRegisterInvalidRequestResponseBody runs the validations defined on
// Register_Invalid_Request_Response_Body
func ValidateRegisterInvalidRequestResponseBody(body *RegisterInvalidRequestResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLoginInvalidRequestResponseBody runs the validations defined on
// Login_Invalid_Request_Response_Body
func ValidateLoginInvalidRequestResponseBody(body *LoginInvalidRequestResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLoginInvalidGroupIDResponseBody runs the validations defined on
// Login_Invalid_Group_ID_Response_Body
func ValidateLoginInvalidGroupIDResponseBody(body *LoginInvalidGroupIDResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateChangeGroupInvalidGroupIDResponseBody runs the validations defined
// on Change_group_Invalid_Group_ID_Response_Body
func ValidateChangeGroupInvalidGroupIDResponseBody(body *ChangeGroupInvalidGroupIDResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateChangeGroupInvalidRequestResponseBody runs the validations defined
// on Change_group_Invalid_Request_Response_Body
func ValidateChangeGroupInvalidRequestResponseBody(body *ChangeGroupInvalidRequestResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDeleteUserInvalidRequestResponseBody runs the validations defined on
// Delete_user_Invalid_Request_Response_Body
func ValidateDeleteUserInvalidRequestResponseBody(body *DeleteUserInvalidRequestResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSaveDataInvalidGroupIDResponseBody runs the validations defined on
// Save_data_Invalid_Group_ID_Response_Body
func ValidateSaveDataInvalidGroupIDResponseBody(body *SaveDataInvalidGroupIDResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSaveDataInvalidRequestResponseBody runs the validations defined on
// Save_data_Invalid_Request_Response_Body
func ValidateSaveDataInvalidRequestResponseBody(body *SaveDataInvalidRequestResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSaveDataInvalidDataResponseBody runs the validations defined on
// Save_data_Invalid_Data_Response_Body
func ValidateSaveDataInvalidDataResponseBody(body *SaveDataInvalidDataResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateReturnDataListInvalidGroupIDResponseBody runs the validations
// defined on Return_data_list_Invalid_Group_ID_Response_Body
func ValidateReturnDataListInvalidGroupIDResponseBody(body *ReturnDataListInvalidGroupIDResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateReturnDataListInvalidRequestResponseBody runs the validations
// defined on Return_data_list_Invalid_Request_Response_Body
func ValidateReturnDataListInvalidRequestResponseBody(body *ReturnDataListInvalidRequestResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePickUpDataInvalidGroupIDResponseBody runs the validations defined on
// Pick_up_data_Invalid_Group_ID_Response_Body
func ValidatePickUpDataInvalidGroupIDResponseBody(body *PickUpDataInvalidGroupIDResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePickUpDataInvalidRequestResponseBody runs the validations defined on
// Pick_up_data_Invalid_Request_Response_Body
func ValidatePickUpDataInvalidRequestResponseBody(body *PickUpDataInvalidRequestResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePickUpDataInvalidDataNameResponseBody runs the validations defined
// on Pick_up_data_Invalid_data_name_Response_Body
func ValidatePickUpDataInvalidDataNameResponseBody(body *PickUpDataInvalidDataNameResponseBody) (err error) {
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}
