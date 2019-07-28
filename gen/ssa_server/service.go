// Code generated by goa v3.0.3, DO NOT EDIT.
//
// SSAServer service
//
// Command:
// $ goa gen SSAServer/design

package ssaserver

import (
	ssaserverviews "SSAServer/gen/ssa_server/views"
	"context"
)

// SSAアプリケーションのサーバーサイド
type Service interface {
	// SSAへの新規登録
	Register(context.Context, *RegisterPayload) (res *SsaResult, err error)
	// SSAへのログイン
	Login(context.Context, *LoginPayload) (res bool, err error)
	// グループIDを変更する
	ChangeGroup(context.Context, *ChangeGroupPayload) (res bool, err error)
	// 既存ユーザーの消去
	DeleteUser(context.Context, *DeleteUserPayload) (res bool, err error)
	// データをサーバーへ保存する
	SaveData(context.Context, *SaveDataPayload) (res bool, err error)
	// データのリストを取得する
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "extended"
	//	- "data"
	//	- "data_extended"
	//	- "data_extended_with_image"
	//	- "data_list_origin"
	ReturnDataList(context.Context, *ReturnDataListPayload) (res SsaResultCollection, view string, err error)
	// データをサーバーから取得する
	PickUpData(context.Context, *PickUpDataPayload) (res *SsaResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "SSAServer"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [7]string{"Register", "Login", "Change_group", "Delete_user", "Save_data", "Return_data_list", "Pick_up_data"}

// RegisterPayload is the payload type of the SSAServer service Register method.
type RegisterPayload struct {
	// User Name
	UserName string
	// User Password
	Password string
	// User mail-address
	Mail string
	// Group ID
	GroupID *string
}

// SsaResult is the result type of the SSAServer service Register method.
type SsaResult struct {
	// User id
	UserID *int
	// User Name
	UserName *string
	// User Password
	Password *string
	// User mail-address
	Mail *string
	// Group ID
	GroupID *string
	// Data name
	DataName *string
	// Data
	Data interface{}
	// Image data
	Image interface{}
	// Data's name
	DataType *int
	// Data title
	Title *string
	// image's name
	ImageName *string
	// date time
	DateTime *string
}

// LoginPayload is the payload type of the SSAServer service Login method.
type LoginPayload struct {
	// User mail-address
	Mail string
	// User Password
	Password string
}

// ChangeGroupPayload is the payload type of the SSAServer service Change_group
// method.
type ChangeGroupPayload struct {
	// User ID
	UserID int
	// User password
	Password string
	// Group ID
	GroupID string
}

// DeleteUserPayload is the payload type of the SSAServer service Delete_user
// method.
type DeleteUserPayload struct {
	// User ID
	UserID int
	// User Password
	Password string
}

// SaveDataPayload is the payload type of the SSAServer service Save_data
// method.
type SaveDataPayload struct {
	// Group ID
	GroupID string
	// User ID
	UserID int
	// Data name
	DataName string
	// Data name
	DataType int
	// Data
	Data []byte
	// Diary title
	Title *string
	// Image name
	ImageName *string
	// Image
	Image []byte
}

// ReturnDataListPayload is the payload type of the SSAServer service
// Return_data_list method.
type ReturnDataListPayload struct {
	// Group ID
	GroupID string
	// User ID
	UserID int
}

// SsaResultCollection is the result type of the SSAServer service
// Return_data_list method.
type SsaResultCollection []*SsaResult

// PickUpDataPayload is the payload type of the SSAServer service Pick_up_data
// method.
type PickUpDataPayload struct {
	// Group ID
	GroupID string
	// Data type
	DataType *string
	// User ID
	UserID int
	// Data name
	DataName string
}

type SsaError struct {
	// Return false when Error occur
	Fault bool
	// Error returned.
	Message string
}

// Error returns an error description.
func (e *SsaError) Error() string {
	return ""
}

// ErrorName returns "SsaError".
func (e *SsaError) ErrorName() string {
	return e.Message
}

// NewSsaResult initializes result type SsaResult from viewed result type
// SsaResult.
func NewSsaResult(vres *ssaserverviews.SsaResult) *SsaResult {
	var res *SsaResult
	switch vres.View {
	case "default", "":
		res = newSsaResult(vres.Projected)
	case "extended":
		res = newSsaResultExtended(vres.Projected)
	case "data":
		res = newSsaResultData(vres.Projected)
	case "data_extended":
		res = newSsaResultDataExtended(vres.Projected)
	case "data_extended_with_image":
		res = newSsaResultDataExtendedWithImage(vres.Projected)
	case "data_list_origin":
		res = newSsaResultDataListOrigin(vres.Projected)
	}
	return res
}

// NewViewedSsaResult initializes viewed result type SsaResult from result type
// SsaResult using the given view.
func NewViewedSsaResult(res *SsaResult, view string) *ssaserverviews.SsaResult {
	var vres *ssaserverviews.SsaResult
	switch view {
	case "default", "":
		p := newSsaResultView(res)
		vres = &ssaserverviews.SsaResult{p, "default"}
	case "extended":
		p := newSsaResultViewExtended(res)
		vres = &ssaserverviews.SsaResult{p, "extended"}
	case "data":
		p := newSsaResultViewData(res)
		vres = &ssaserverviews.SsaResult{p, "data"}
	case "data_extended":
		p := newSsaResultViewDataExtended(res)
		vres = &ssaserverviews.SsaResult{p, "data_extended"}
	case "data_extended_with_image":
		p := newSsaResultViewDataExtendedWithImage(res)
		vres = &ssaserverviews.SsaResult{p, "data_extended_with_image"}
	case "data_list_origin":
		p := newSsaResultViewDataListOrigin(res)
		vres = &ssaserverviews.SsaResult{p, "data_list_origin"}
	}
	return vres
}

// NewSsaResultCollection initializes result type SsaResultCollection from
// viewed result type SsaResultCollection.
func NewSsaResultCollection(vres ssaserverviews.SsaResultCollection) SsaResultCollection {
	var res SsaResultCollection
	switch vres.View {
	case "default", "":
		res = newSsaResultCollection(vres.Projected)
	case "extended":
		res = newSsaResultCollectionExtended(vres.Projected)
	case "data":
		res = newSsaResultCollectionData(vres.Projected)
	case "data_extended":
		res = newSsaResultCollectionDataExtended(vres.Projected)
	case "data_extended_with_image":
		res = newSsaResultCollectionDataExtendedWithImage(vres.Projected)
	case "data_list_origin":
		res = newSsaResultCollectionDataListOrigin(vres.Projected)
	}
	return res
}

// NewViewedSsaResultCollection initializes viewed result type
// SsaResultCollection from result type SsaResultCollection using the given
// view.
func NewViewedSsaResultCollection(res SsaResultCollection, view string) ssaserverviews.SsaResultCollection {
	var vres ssaserverviews.SsaResultCollection
	switch view {
	case "default", "":
		p := newSsaResultCollectionView(res)
		vres = ssaserverviews.SsaResultCollection{p, "default"}
	case "extended":
		p := newSsaResultCollectionViewExtended(res)
		vres = ssaserverviews.SsaResultCollection{p, "extended"}
	case "data":
		p := newSsaResultCollectionViewData(res)
		vres = ssaserverviews.SsaResultCollection{p, "data"}
	case "data_extended":
		p := newSsaResultCollectionViewDataExtended(res)
		vres = ssaserverviews.SsaResultCollection{p, "data_extended"}
	case "data_extended_with_image":
		p := newSsaResultCollectionViewDataExtendedWithImage(res)
		vres = ssaserverviews.SsaResultCollection{p, "data_extended_with_image"}
	case "data_list_origin":
		p := newSsaResultCollectionViewDataListOrigin(res)
		vres = ssaserverviews.SsaResultCollection{p, "data_list_origin"}
	}
	return vres
}

// newSsaResult converts projected type SsaResult to service type SsaResult.
func newSsaResult(vres *ssaserverviews.SsaResultView) *SsaResult {
	res := &SsaResult{
		UserName: vres.UserName,
	}
	if vres.UserID != nil {
		res.UserID = vres.UserID
	}
	if vres.Mail != nil {
		res.Mail = vres.Mail
	}
	return res
}

// newSsaResultExtended converts projected type SsaResult to service type
// SsaResult.
func newSsaResultExtended(vres *ssaserverviews.SsaResultView) *SsaResult {
	res := &SsaResult{
		UserName: vres.UserName,
		GroupID:  vres.GroupID,
	}
	if vres.UserID != nil {
		res.UserID = vres.UserID
	}
	if vres.Mail != nil {
		res.Mail = vres.Mail
	}
	return res
}

// newSsaResultData converts projected type SsaResult to service type SsaResult.
func newSsaResultData(vres *ssaserverviews.SsaResultView) *SsaResult {
	res := &SsaResult{
		Data:     vres.Data,
		DataType: vres.DataType,
		DataName: vres.DataName,
	}
	return res
}

// newSsaResultDataExtended converts projected type SsaResult to service type
// SsaResult.
func newSsaResultDataExtended(vres *ssaserverviews.SsaResultView) *SsaResult {
	res := &SsaResult{
		Data:     vres.Data,
		DataType: vres.DataType,
		DataName: vres.DataName,
		Title:    vres.Title,
	}
	return res
}

// newSsaResultDataExtendedWithImage converts projected type SsaResult to
// service type SsaResult.
func newSsaResultDataExtendedWithImage(vres *ssaserverviews.SsaResultView) *SsaResult {
	res := &SsaResult{
		Data:      vres.Data,
		DataType:  vres.DataType,
		DataName:  vres.DataName,
		Title:     vres.Title,
		Image:     vres.Image,
		ImageName: vres.ImageName,
	}
	return res
}

// newSsaResultDataListOrigin converts projected type SsaResult to service type
// SsaResult.
func newSsaResultDataListOrigin(vres *ssaserverviews.SsaResultView) *SsaResult {
	res := &SsaResult{
		DataType: vres.DataType,
		DataName: vres.DataName,
		Title:    vres.Title,
		DateTime: vres.DateTime,
		UserName: vres.UserName,
	}
	return res
}

// newSsaResultView projects result type SsaResult to projected type
// SsaResultView using the "default" view.
func newSsaResultView(res *SsaResult) *ssaserverviews.SsaResultView {
	vres := &ssaserverviews.SsaResultView{
		UserID:   res.UserID,
		UserName: res.UserName,
		Mail:     res.Mail,
	}
	return vres
}

// newSsaResultViewExtended projects result type SsaResult to projected type
// SsaResultView using the "extended" view.
func newSsaResultViewExtended(res *SsaResult) *ssaserverviews.SsaResultView {
	vres := &ssaserverviews.SsaResultView{
		UserID:   res.UserID,
		UserName: res.UserName,
		Mail:     res.Mail,
		GroupID:  res.GroupID,
	}
	return vres
}

// newSsaResultViewData projects result type SsaResult to projected type
// SsaResultView using the "data" view.
func newSsaResultViewData(res *SsaResult) *ssaserverviews.SsaResultView {
	vres := &ssaserverviews.SsaResultView{
		DataName: res.DataName,
		Data:     res.Data,
		DataType: res.DataType,
	}
	return vres
}

// newSsaResultViewDataExtended projects result type SsaResult to projected
// type SsaResultView using the "data_extended" view.
func newSsaResultViewDataExtended(res *SsaResult) *ssaserverviews.SsaResultView {
	vres := &ssaserverviews.SsaResultView{
		DataName: res.DataName,
		Data:     res.Data,
		DataType: res.DataType,
		Title:    res.Title,
	}
	return vres
}

// newSsaResultViewDataExtendedWithImage projects result type SsaResult to
// projected type SsaResultView using the "data_extended_with_image" view.
func newSsaResultViewDataExtendedWithImage(res *SsaResult) *ssaserverviews.SsaResultView {
	vres := &ssaserverviews.SsaResultView{
		DataName:  res.DataName,
		Data:      res.Data,
		Image:     res.Image,
		DataType:  res.DataType,
		Title:     res.Title,
		ImageName: res.ImageName,
	}
	return vres
}

// newSsaResultViewDataListOrigin projects result type SsaResult to projected
// type SsaResultView using the "data_list_origin" view.
func newSsaResultViewDataListOrigin(res *SsaResult) *ssaserverviews.SsaResultView {
	vres := &ssaserverviews.SsaResultView{
		UserName: res.UserName,
		DataName: res.DataName,
		DataType: res.DataType,
		Title:    res.Title,
		DateTime: res.DateTime,
	}
	return vres
}

// newSsaResultCollection converts projected type SsaResultCollection to
// service type SsaResultCollection.
func newSsaResultCollection(vres ssaserverviews.SsaResultCollectionView) SsaResultCollection {
	res := make(SsaResultCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsaResult(n)
	}
	return res
}

// newSsaResultCollectionExtended converts projected type SsaResultCollection
// to service type SsaResultCollection.
func newSsaResultCollectionExtended(vres ssaserverviews.SsaResultCollectionView) SsaResultCollection {
	res := make(SsaResultCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsaResultExtended(n)
	}
	return res
}

// newSsaResultCollectionData converts projected type SsaResultCollection to
// service type SsaResultCollection.
func newSsaResultCollectionData(vres ssaserverviews.SsaResultCollectionView) SsaResultCollection {
	res := make(SsaResultCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsaResultData(n)
	}
	return res
}

// newSsaResultCollectionDataExtended converts projected type
// SsaResultCollection to service type SsaResultCollection.
func newSsaResultCollectionDataExtended(vres ssaserverviews.SsaResultCollectionView) SsaResultCollection {
	res := make(SsaResultCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsaResultDataExtended(n)
	}
	return res
}

// newSsaResultCollectionDataExtendedWithImage converts projected type
// SsaResultCollection to service type SsaResultCollection.
func newSsaResultCollectionDataExtendedWithImage(vres ssaserverviews.SsaResultCollectionView) SsaResultCollection {
	res := make(SsaResultCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsaResultDataExtendedWithImage(n)
	}
	return res
}

// newSsaResultCollectionDataListOrigin converts projected type
// SsaResultCollection to service type SsaResultCollection.
func newSsaResultCollectionDataListOrigin(vres ssaserverviews.SsaResultCollectionView) SsaResultCollection {
	res := make(SsaResultCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsaResultDataListOrigin(n)
	}
	return res
}

// newSsaResultCollectionView projects result type SsaResultCollection to
// projected type SsaResultCollectionView using the "default" view.
func newSsaResultCollectionView(res SsaResultCollection) ssaserverviews.SsaResultCollectionView {
	vres := make(ssaserverviews.SsaResultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsaResultView(n)
	}
	return vres
}

// newSsaResultCollectionViewExtended projects result type SsaResultCollection
// to projected type SsaResultCollectionView using the "extended" view.
func newSsaResultCollectionViewExtended(res SsaResultCollection) ssaserverviews.SsaResultCollectionView {
	vres := make(ssaserverviews.SsaResultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsaResultViewExtended(n)
	}
	return vres
}

// newSsaResultCollectionViewData projects result type SsaResultCollection to
// projected type SsaResultCollectionView using the "data" view.
func newSsaResultCollectionViewData(res SsaResultCollection) ssaserverviews.SsaResultCollectionView {
	vres := make(ssaserverviews.SsaResultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsaResultViewData(n)
	}
	return vres
}

// newSsaResultCollectionViewDataExtended projects result type
// SsaResultCollection to projected type SsaResultCollectionView using the
// "data_extended" view.
func newSsaResultCollectionViewDataExtended(res SsaResultCollection) ssaserverviews.SsaResultCollectionView {
	vres := make(ssaserverviews.SsaResultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsaResultViewDataExtended(n)
	}
	return vres
}

// newSsaResultCollectionViewDataExtendedWithImage projects result type
// SsaResultCollection to projected type SsaResultCollectionView using the
// "data_extended_with_image" view.
func newSsaResultCollectionViewDataExtendedWithImage(res SsaResultCollection) ssaserverviews.SsaResultCollectionView {
	vres := make(ssaserverviews.SsaResultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsaResultViewDataExtendedWithImage(n)
	}
	return vres
}

// newSsaResultCollectionViewDataListOrigin projects result type
// SsaResultCollection to projected type SsaResultCollectionView using the
// "data_list_origin" view.
func newSsaResultCollectionViewDataListOrigin(res SsaResultCollection) ssaserverviews.SsaResultCollectionView {
	vres := make(ssaserverviews.SsaResultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsaResultViewDataListOrigin(n)
	}
	return vres
}
