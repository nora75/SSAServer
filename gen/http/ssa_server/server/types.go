// Code generated by goa v3.0.3, DO NOT EDIT.
//
// SSAServer HTTP server types
//
// Command:
// $ goa gen SSAServer/design

package server

import (
	ssaserver "SSAServer/gen/ssa_server"
	ssaserverviews "SSAServer/gen/ssa_server/views"

	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "SSAServer" service "Register"
// endpoint HTTP request body.
type RegisterRequestBody struct {
	// User Name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// User Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// User mail-address
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// Group ID
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" xml:"group_id,omitempty"`
}

// LoginRequestBody is the type of the "SSAServer" service "Login" endpoint
// HTTP request body.
type LoginRequestBody struct {
	// User mail-address
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// User Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// ChangeGroupRequestBody is the type of the "SSAServer" service "Change_group"
// endpoint HTTP request body.
type ChangeGroupRequestBody struct {
	// User password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Group ID
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" xml:"group_id,omitempty"`
}

// DeleteUserRequestBody is the type of the "SSAServer" service "Delete_user"
// endpoint HTTP request body.
type DeleteUserRequestBody struct {
	// User Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// SaveDataRequestBody is the type of the "SSAServer" service "Save_data"
// endpoint HTTP request body.
type SaveDataRequestBody struct {
	// User ID
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// Data name
	DataName *string `form:"data_name,omitempty" json:"data_name,omitempty" xml:"data_name,omitempty"`
	// Data name
	DataType *int `form:"data_type,omitempty" json:"data_type,omitempty" xml:"data_type,omitempty"`
	// Data
	Data []byte `form:"Data,omitempty" json:"Data,omitempty" xml:"Data,omitempty"`
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
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// PickUpDataRequestBody is the type of the "SSAServer" service "Pick_up_data"
// endpoint HTTP request body.
type PickUpDataRequestBody struct {
	// User ID
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// Data name
	DataName *string `form:"data_name,omitempty" json:"data_name,omitempty" xml:"data_name,omitempty"`
	// Data name
	IamgeName *string `form:"iamge_name,omitempty" json:"iamge_name,omitempty" xml:"iamge_name,omitempty"`
	// Data's User ID
	DataUserID *int `form:"data_user_id,omitempty" json:"data_user_id,omitempty" xml:"data_user_id,omitempty"`
}

// RegisterResponseBodyExtended is the type of the "SSAServer" service
// "Register" endpoint HTTP response body.
type RegisterResponseBodyExtended struct {
	// User id
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
	// User mail-address
	Mail string `form:"mail" json:"mail" xml:"mail"`
	// User Name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// Group ID
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" xml:"group_id,omitempty"`
}

// SsaResultResponseCollection is the type of the "SSAServer" service
// "Return_data_list" endpoint HTTP response body.
type SsaResultResponseCollection []*SsaResultResponse

// SsaResultResponseExtendedCollection is the type of the "SSAServer" service
// "Return_data_list" endpoint HTTP response body.
type SsaResultResponseExtendedCollection []*SsaResultResponseExtended

// SsaResultResponseDataCollection is the type of the "SSAServer" service
// "Return_data_list" endpoint HTTP response body.
type SsaResultResponseDataCollection []*SsaResultResponseData

// SsaResultResponseDataExtendedCollection is the type of the "SSAServer"
// service "Return_data_list" endpoint HTTP response body.
type SsaResultResponseDataExtendedCollection []*SsaResultResponseDataExtended

// SsaResultResponseDataExtendedWithImageCollection is the type of the
// "SSAServer" service "Return_data_list" endpoint HTTP response body.
type SsaResultResponseDataExtendedWithImageCollection []*SsaResultResponseDataExtendedWithImage

// SsaResultResponseDataListOriginCollection is the type of the "SSAServer"
// service "Return_data_list" endpoint HTTP response body.
type SsaResultResponseDataListOriginCollection []*SsaResultResponseDataListOrigin

// PickUpDataResponseBodyDataExtendedWithImage is the type of the "SSAServer"
// service "Pick_up_data" endpoint HTTP response body.
type PickUpDataResponseBodyDataExtendedWithImage struct {
	// Data
	Data []byte `form:"Data" json:"Data" xml:"Data"`
	// Data's name
	DataType int `form:"data_type" json:"data_type" xml:"data_type"`
	// Data name
	DataName string `form:"data_name" json:"data_name" xml:"data_name"`
	// Data title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Image data
	Image []byte `form:"Image,omitempty" json:"Image,omitempty" xml:"Image,omitempty"`
	// image's name
	ImageName *string `form:"image_name,omitempty" json:"image_name,omitempty" xml:"image_name,omitempty"`
}

// RegisterInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Register" endpoint HTTP response body for the "Invalid_Group_ID" error.
type RegisterInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// RegisterTheUserAlreadyExistsResponseBody is the type of the "SSAServer"
// service "Register" endpoint HTTP response body for the
// "The_user_already_exists" error.
type RegisterTheUserAlreadyExistsResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// RegisterInvalidRequestResponseBody is the type of the "SSAServer" service
// "Register" endpoint HTTP response body for the "Invalid_Request" error.
type RegisterInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// LoginInvalidRequestResponseBody is the type of the "SSAServer" service
// "Login" endpoint HTTP response body for the "Invalid_Request" error.
type LoginInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// LoginInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Login" endpoint HTTP response body for the "Invalid_Group_ID" error.
type LoginInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// ChangeGroupInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Change_group" endpoint HTTP response body for the "Invalid_Group_ID" error.
type ChangeGroupInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// ChangeGroupInvalidRequestResponseBody is the type of the "SSAServer" service
// "Change_group" endpoint HTTP response body for the "Invalid_Request" error.
type ChangeGroupInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteUserInvalidRequestResponseBody is the type of the "SSAServer" service
// "Delete_user" endpoint HTTP response body for the "Invalid_Request" error.
type DeleteUserInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// SaveDataInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Save_data" endpoint HTTP response body for the "Invalid_Group_ID" error.
type SaveDataInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// SaveDataInvalidRequestResponseBody is the type of the "SSAServer" service
// "Save_data" endpoint HTTP response body for the "Invalid_Request" error.
type SaveDataInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// SaveDataInvalidDataResponseBody is the type of the "SSAServer" service
// "Save_data" endpoint HTTP response body for the "Invalid_Data" error.
type SaveDataInvalidDataResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// ReturnDataListInvalidGroupIDResponseBody is the type of the "SSAServer"
// service "Return_data_list" endpoint HTTP response body for the
// "Invalid_Group_ID" error.
type ReturnDataListInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// ReturnDataListInvalidRequestResponseBody is the type of the "SSAServer"
// service "Return_data_list" endpoint HTTP response body for the
// "Invalid_Request" error.
type ReturnDataListInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// PickUpDataInvalidGroupIDResponseBody is the type of the "SSAServer" service
// "Pick_up_data" endpoint HTTP response body for the "Invalid_Group_ID" error.
type PickUpDataInvalidGroupIDResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// PickUpDataInvalidRequestResponseBody is the type of the "SSAServer" service
// "Pick_up_data" endpoint HTTP response body for the "Invalid_Request" error.
type PickUpDataInvalidRequestResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// PickUpDataInvalidDataNameResponseBody is the type of the "SSAServer" service
// "Pick_up_data" endpoint HTTP response body for the "Invalid_data_name" error.
type PickUpDataInvalidDataNameResponseBody struct {
	// Return false when Error occur
	Fault bool `form:"fault" json:"fault" xml:"fault"`
	// Error returned.
	Message string `form:"message" json:"message" xml:"message"`
}

// SsaResultResponse is used to define fields on response body types.
type SsaResultResponse struct {
	// User id
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User mail-address
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// User Name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
}

// SsaResultResponseExtended is used to define fields on response body types.
type SsaResultResponseExtended struct {
	// User id
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User mail-address
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// User Name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// Group ID
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" xml:"group_id,omitempty"`
}

// SsaResultResponseData is used to define fields on response body types.
type SsaResultResponseData struct {
	// Data
	Data []byte `form:"Data,omitempty" json:"Data,omitempty" xml:"Data,omitempty"`
	// Data's name
	DataType *int `form:"data_type,omitempty" json:"data_type,omitempty" xml:"data_type,omitempty"`
	// Data name
	DataName *string `form:"data_name,omitempty" json:"data_name,omitempty" xml:"data_name,omitempty"`
}

// SsaResultResponseDataExtended is used to define fields on response body
// types.
type SsaResultResponseDataExtended struct {
	// Data
	Data []byte `form:"Data,omitempty" json:"Data,omitempty" xml:"Data,omitempty"`
	// Data's name
	DataType *int `form:"data_type,omitempty" json:"data_type,omitempty" xml:"data_type,omitempty"`
	// Data name
	DataName *string `form:"data_name,omitempty" json:"data_name,omitempty" xml:"data_name,omitempty"`
	// Data title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// SsaResultResponseDataExtendedWithImage is used to define fields on response
// body types.
type SsaResultResponseDataExtendedWithImage struct {
	// Data
	Data []byte `form:"Data,omitempty" json:"Data,omitempty" xml:"Data,omitempty"`
	// Data's name
	DataType *int `form:"data_type,omitempty" json:"data_type,omitempty" xml:"data_type,omitempty"`
	// Data name
	DataName *string `form:"data_name,omitempty" json:"data_name,omitempty" xml:"data_name,omitempty"`
	// Data title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Image data
	Image []byte `form:"Image,omitempty" json:"Image,omitempty" xml:"Image,omitempty"`
	// image's name
	ImageName *string `form:"image_name,omitempty" json:"image_name,omitempty" xml:"image_name,omitempty"`
}

// SsaResultResponseDataListOrigin is used to define fields on response body
// types.
type SsaResultResponseDataListOrigin struct {
	// Data's name
	DataType *int `form:"data_type,omitempty" json:"data_type,omitempty" xml:"data_type,omitempty"`
	// Data name
	DataName *string `form:"data_name,omitempty" json:"data_name,omitempty" xml:"data_name,omitempty"`
	// image's name
	ImageName *string `form:"image_name,omitempty" json:"image_name,omitempty" xml:"image_name,omitempty"`
	// Data title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// date time
	DateTime *string `form:"date_time,omitempty" json:"date_time,omitempty" xml:"date_time,omitempty"`
	// User Name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// User id
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// NewRegisterResponseBodyExtended builds the HTTP response body from the
// result of the "Register" endpoint of the "SSAServer" service.
func NewRegisterResponseBodyExtended(res *ssaserverviews.SsaResultView) *RegisterResponseBodyExtended {
	body := &RegisterResponseBodyExtended{
		UserName: res.UserName,
		GroupID:  res.GroupID,
	}
	if res.UserID != nil {
		body.UserID = *res.UserID
	}
	if res.Mail != nil {
		body.Mail = *res.Mail
	}
	return body
}

// NewSsaResultResponseCollection builds the HTTP response body from the result
// of the "Return_data_list" endpoint of the "SSAServer" service.
func NewSsaResultResponseCollection(res ssaserverviews.SsaResultCollectionView) SsaResultResponseCollection {
	body := make([]*SsaResultResponse, len(res))
	for i, val := range res {
		body[i] = &SsaResultResponse{
			UserID:   val.UserID,
			UserName: val.UserName,
			Mail:     val.Mail,
		}
	}
	return body
}

// NewSsaResultResponseExtendedCollection builds the HTTP response body from
// the result of the "Return_data_list" endpoint of the "SSAServer" service.
func NewSsaResultResponseExtendedCollection(res ssaserverviews.SsaResultCollectionView) SsaResultResponseExtendedCollection {
	body := make([]*SsaResultResponseExtended, len(res))
	for i, val := range res {
		body[i] = &SsaResultResponseExtended{
			UserID:   val.UserID,
			UserName: val.UserName,
			Mail:     val.Mail,
			GroupID:  val.GroupID,
		}
	}
	return body
}

// NewSsaResultResponseDataCollection builds the HTTP response body from the
// result of the "Return_data_list" endpoint of the "SSAServer" service.
func NewSsaResultResponseDataCollection(res ssaserverviews.SsaResultCollectionView) SsaResultResponseDataCollection {
	body := make([]*SsaResultResponseData, len(res))
	for i, val := range res {
		body[i] = &SsaResultResponseData{
			DataName: val.DataName,
			Data:     val.Data,
			DataType: val.DataType,
		}
	}
	return body
}

// NewSsaResultResponseDataExtendedCollection builds the HTTP response body
// from the result of the "Return_data_list" endpoint of the "SSAServer"
// service.
func NewSsaResultResponseDataExtendedCollection(res ssaserverviews.SsaResultCollectionView) SsaResultResponseDataExtendedCollection {
	body := make([]*SsaResultResponseDataExtended, len(res))
	for i, val := range res {
		body[i] = &SsaResultResponseDataExtended{
			DataName: val.DataName,
			Data:     val.Data,
			DataType: val.DataType,
			Title:    val.Title,
		}
	}
	return body
}

// NewSsaResultResponseDataExtendedWithImageCollection builds the HTTP response
// body from the result of the "Return_data_list" endpoint of the "SSAServer"
// service.
func NewSsaResultResponseDataExtendedWithImageCollection(res ssaserverviews.SsaResultCollectionView) SsaResultResponseDataExtendedWithImageCollection {
	body := make([]*SsaResultResponseDataExtendedWithImage, len(res))
	for i, val := range res {
		body[i] = &SsaResultResponseDataExtendedWithImage{
			DataName:  val.DataName,
			Data:      val.Data,
			Image:     val.Image,
			DataType:  val.DataType,
			Title:     val.Title,
			ImageName: val.ImageName,
		}
	}
	return body
}

// NewSsaResultResponseDataListOriginCollection builds the HTTP response body
// from the result of the "Return_data_list" endpoint of the "SSAServer"
// service.
func NewSsaResultResponseDataListOriginCollection(res ssaserverviews.SsaResultCollectionView) SsaResultResponseDataListOriginCollection {
	body := make([]*SsaResultResponseDataListOrigin, len(res))
	for i, val := range res {
		body[i] = &SsaResultResponseDataListOrigin{
			UserID:    val.UserID,
			UserName:  val.UserName,
			DataName:  val.DataName,
			DataType:  val.DataType,
			Title:     val.Title,
			ImageName: val.ImageName,
			DateTime:  val.DateTime,
		}
	}
	return body
}

// NewPickUpDataResponseBodyDataExtendedWithImage builds the HTTP response body
// from the result of the "Pick_up_data" endpoint of the "SSAServer" service.
func NewPickUpDataResponseBodyDataExtendedWithImage(res *ssaserverviews.SsaResultView) *PickUpDataResponseBodyDataExtendedWithImage {
	body := &PickUpDataResponseBodyDataExtendedWithImage{
		Data:      res.Data,
		Image:     res.Image,
		Title:     res.Title,
		ImageName: res.ImageName,
	}
	if res.DataName != nil {
		body.DataName = *res.DataName
	}
	if res.DataType != nil {
		body.DataType = *res.DataType
	}
	return body
}

// NewRegisterInvalidGroupIDResponseBody builds the HTTP response body from the
// result of the "Register" endpoint of the "SSAServer" service.
func NewRegisterInvalidGroupIDResponseBody(res *ssaserver.SsaError) *RegisterInvalidGroupIDResponseBody {
	body := &RegisterInvalidGroupIDResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewRegisterTheUserAlreadyExistsResponseBody builds the HTTP response body
// from the result of the "Register" endpoint of the "SSAServer" service.
func NewRegisterTheUserAlreadyExistsResponseBody(res *ssaserver.SsaError) *RegisterTheUserAlreadyExistsResponseBody {
	body := &RegisterTheUserAlreadyExistsResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewRegisterInvalidRequestResponseBody builds the HTTP response body from the
// result of the "Register" endpoint of the "SSAServer" service.
func NewRegisterInvalidRequestResponseBody(res *ssaserver.SsaError) *RegisterInvalidRequestResponseBody {
	body := &RegisterInvalidRequestResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewLoginInvalidRequestResponseBody builds the HTTP response body from the
// result of the "Login" endpoint of the "SSAServer" service.
func NewLoginInvalidRequestResponseBody(res *ssaserver.SsaError) *LoginInvalidRequestResponseBody {
	body := &LoginInvalidRequestResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewLoginInvalidGroupIDResponseBody builds the HTTP response body from the
// result of the "Login" endpoint of the "SSAServer" service.
func NewLoginInvalidGroupIDResponseBody(res *ssaserver.SsaError) *LoginInvalidGroupIDResponseBody {
	body := &LoginInvalidGroupIDResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewChangeGroupInvalidGroupIDResponseBody builds the HTTP response body from
// the result of the "Change_group" endpoint of the "SSAServer" service.
func NewChangeGroupInvalidGroupIDResponseBody(res *ssaserver.SsaError) *ChangeGroupInvalidGroupIDResponseBody {
	body := &ChangeGroupInvalidGroupIDResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewChangeGroupInvalidRequestResponseBody builds the HTTP response body from
// the result of the "Change_group" endpoint of the "SSAServer" service.
func NewChangeGroupInvalidRequestResponseBody(res *ssaserver.SsaError) *ChangeGroupInvalidRequestResponseBody {
	body := &ChangeGroupInvalidRequestResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewDeleteUserInvalidRequestResponseBody builds the HTTP response body from
// the result of the "Delete_user" endpoint of the "SSAServer" service.
func NewDeleteUserInvalidRequestResponseBody(res *ssaserver.SsaError) *DeleteUserInvalidRequestResponseBody {
	body := &DeleteUserInvalidRequestResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewSaveDataInvalidGroupIDResponseBody builds the HTTP response body from the
// result of the "Save_data" endpoint of the "SSAServer" service.
func NewSaveDataInvalidGroupIDResponseBody(res *ssaserver.SsaError) *SaveDataInvalidGroupIDResponseBody {
	body := &SaveDataInvalidGroupIDResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewSaveDataInvalidRequestResponseBody builds the HTTP response body from the
// result of the "Save_data" endpoint of the "SSAServer" service.
func NewSaveDataInvalidRequestResponseBody(res *ssaserver.SsaError) *SaveDataInvalidRequestResponseBody {
	body := &SaveDataInvalidRequestResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewSaveDataInvalidDataResponseBody builds the HTTP response body from the
// result of the "Save_data" endpoint of the "SSAServer" service.
func NewSaveDataInvalidDataResponseBody(res *ssaserver.SsaError) *SaveDataInvalidDataResponseBody {
	body := &SaveDataInvalidDataResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewReturnDataListInvalidGroupIDResponseBody builds the HTTP response body
// from the result of the "Return_data_list" endpoint of the "SSAServer"
// service.
func NewReturnDataListInvalidGroupIDResponseBody(res *ssaserver.SsaError) *ReturnDataListInvalidGroupIDResponseBody {
	body := &ReturnDataListInvalidGroupIDResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewReturnDataListInvalidRequestResponseBody builds the HTTP response body
// from the result of the "Return_data_list" endpoint of the "SSAServer"
// service.
func NewReturnDataListInvalidRequestResponseBody(res *ssaserver.SsaError) *ReturnDataListInvalidRequestResponseBody {
	body := &ReturnDataListInvalidRequestResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewPickUpDataInvalidGroupIDResponseBody builds the HTTP response body from
// the result of the "Pick_up_data" endpoint of the "SSAServer" service.
func NewPickUpDataInvalidGroupIDResponseBody(res *ssaserver.SsaError) *PickUpDataInvalidGroupIDResponseBody {
	body := &PickUpDataInvalidGroupIDResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewPickUpDataInvalidRequestResponseBody builds the HTTP response body from
// the result of the "Pick_up_data" endpoint of the "SSAServer" service.
func NewPickUpDataInvalidRequestResponseBody(res *ssaserver.SsaError) *PickUpDataInvalidRequestResponseBody {
	body := &PickUpDataInvalidRequestResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewPickUpDataInvalidDataNameResponseBody builds the HTTP response body from
// the result of the "Pick_up_data" endpoint of the "SSAServer" service.
func NewPickUpDataInvalidDataNameResponseBody(res *ssaserver.SsaError) *PickUpDataInvalidDataNameResponseBody {
	body := &PickUpDataInvalidDataNameResponseBody{
		Fault:   res.Fault,
		Message: res.Message,
	}
	return body
}

// NewRegisterPayload builds a SSAServer service Register endpoint payload.
func NewRegisterPayload(body *RegisterRequestBody) *ssaserver.RegisterPayload {
	v := &ssaserver.RegisterPayload{
		UserName: *body.UserName,
		Password: *body.Password,
		Mail:     *body.Mail,
		GroupID:  body.GroupID,
	}
	return v
}

// NewLoginPayload builds a SSAServer service Login endpoint payload.
func NewLoginPayload(body *LoginRequestBody) *ssaserver.LoginPayload {
	v := &ssaserver.LoginPayload{
		Mail:     *body.Mail,
		Password: *body.Password,
	}
	return v
}

// NewChangeGroupPayload builds a SSAServer service Change_group endpoint
// payload.
func NewChangeGroupPayload(body *ChangeGroupRequestBody, userID int) *ssaserver.ChangeGroupPayload {
	v := &ssaserver.ChangeGroupPayload{
		Password: *body.Password,
		GroupID:  *body.GroupID,
	}
	v.UserID = userID
	return v
}

// NewDeleteUserPayload builds a SSAServer service Delete_user endpoint payload.
func NewDeleteUserPayload(body *DeleteUserRequestBody, userID int) *ssaserver.DeleteUserPayload {
	v := &ssaserver.DeleteUserPayload{
		Password: *body.Password,
	}
	v.UserID = userID
	return v
}

// NewSaveDataPayload builds a SSAServer service Save_data endpoint payload.
func NewSaveDataPayload(body *SaveDataRequestBody, groupID string) *ssaserver.SaveDataPayload {
	v := &ssaserver.SaveDataPayload{
		UserID:    *body.UserID,
		DataName:  *body.DataName,
		DataType:  *body.DataType,
		Data:      body.Data,
		Title:     body.Title,
		ImageName: body.ImageName,
		Image:     body.Image,
	}
	v.GroupID = groupID
	return v
}

// NewReturnDataListPayload builds a SSAServer service Return_data_list
// endpoint payload.
func NewReturnDataListPayload(body *ReturnDataListRequestBody, groupID string) *ssaserver.ReturnDataListPayload {
	v := &ssaserver.ReturnDataListPayload{
		UserID: *body.UserID,
	}
	v.GroupID = groupID
	return v
}

// NewPickUpDataPayload builds a SSAServer service Pick_up_data endpoint
// payload.
func NewPickUpDataPayload(body *PickUpDataRequestBody, groupID string, dataType string) *ssaserver.PickUpDataPayload {
	v := &ssaserver.PickUpDataPayload{
		UserID:     *body.UserID,
		DataName:   *body.DataName,
		IamgeName:  body.IamgeName,
		DataUserID: *body.DataUserID,
	}
	v.GroupID = groupID
	v.DataType = &dataType
	return v
}

// ValidateRegisterRequestBody runs the validations defined on
// RegisterRequestBody
func ValidateRegisterRequestBody(body *RegisterRequestBody) (err error) {
	if body.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_name", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.Mail == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mail", "body"))
	}
	if body.Mail != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.mail", *body.Mail, goa.FormatEmail))
	}
	return
}

// ValidateLoginRequestBody runs the validations defined on LoginRequestBody
func ValidateLoginRequestBody(body *LoginRequestBody) (err error) {
	if body.Mail == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mail", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}

// ValidateChangeGroupRequestBody runs the validations defined on
// Change_group_Request_Body
func ValidateChangeGroupRequestBody(body *ChangeGroupRequestBody) (err error) {
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.GroupID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("group_id", "body"))
	}
	return
}

// ValidateDeleteUserRequestBody runs the validations defined on
// Delete_user_Request_Body
func ValidateDeleteUserRequestBody(body *DeleteUserRequestBody) (err error) {
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}

// ValidateSaveDataRequestBody runs the validations defined on
// Save_data_Request_Body
func ValidateSaveDataRequestBody(body *SaveDataRequestBody) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "body"))
	}
	if body.DataName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data_name", "body"))
	}
	if body.DataType == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data_type", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Data", "body"))
	}
	return
}

// ValidateReturnDataListRequestBody runs the validations defined on
// Return_data_list_Request_Body
func ValidateReturnDataListRequestBody(body *ReturnDataListRequestBody) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "body"))
	}
	return
}

// ValidatePickUpDataRequestBody runs the validations defined on
// Pick_up_data_Request_Body
func ValidatePickUpDataRequestBody(body *PickUpDataRequestBody) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "body"))
	}
	if body.DataName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data_name", "body"))
	}
	if body.DataUserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data_user_id", "body"))
	}
	return
}
