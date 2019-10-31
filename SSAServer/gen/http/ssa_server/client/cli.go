// Code generated by goa v3.0.6, DO NOT EDIT.
//
// SSAServer HTTP client CLI support package
//
// Command:
// $ goa gen SSAServer/design

package client

import (
	ssaserver "SSAServer/gen/ssa_server"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildRegisterPayload builds the payload for the SSAServer Register endpoint
// from CLI flags.
func BuildRegisterPayload(sSAServerRegisterBody string) (*ssaserver.RegisterPayload, error) {
	var err error
	var body RegisterRequestBody
	{
		err = json.Unmarshal([]byte(sSAServerRegisterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"group_id\": \"group-bzuiphjjgdas\",\n      \"mail\": \"hoge@hoge.com\",\n      \"password\": \"passW0rd\",\n      \"user_name\": \"平野竜也\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.mail", body.Mail, goa.FormatEmail))

		if err != nil {
			return nil, err
		}
	}
	v := &ssaserver.RegisterPayload{
		UserName: body.UserName,
		Password: body.Password,
		Mail:     body.Mail,
		GroupID:  body.GroupID,
	}
	return v, nil
}

// BuildLoginPayload builds the payload for the SSAServer Login endpoint from
// CLI flags.
func BuildLoginPayload(sSAServerLoginBody string) (*ssaserver.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(sSAServerLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"mail\": \"hoge@hoge.com\",\n      \"password\": \"password\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.mail", body.Mail, goa.FormatEmail))

		if err != nil {
			return nil, err
		}
	}
	v := &ssaserver.LoginPayload{
		Mail:     body.Mail,
		Password: body.Password,
	}
	return v, nil
}

// BuildChangeGroupPayload builds the payload for the SSAServer Change_group
// endpoint from CLI flags.
func BuildChangeGroupPayload(sSAServerChangeGroupBody string, sSAServerChangeGroupUserID string) (*ssaserver.ChangeGroupPayload, error) {
	var err error
	var body ChangeGroupRequestBody
	{
		err = json.Unmarshal([]byte(sSAServerChangeGroupBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"group_id\": \"group2\",\n      \"password\": \"pass1234\"\n   }'")
		}
	}
	var userID int
	{
		var v int64
		v, err = strconv.ParseInt(sSAServerChangeGroupUserID, 10, 64)
		userID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for userID, must be INT")
		}
	}
	v := &ssaserver.ChangeGroupPayload{
		Password: body.Password,
		GroupID:  body.GroupID,
	}
	v.UserID = userID
	return v, nil
}

// BuildDeleteUserPayload builds the payload for the SSAServer Delete_user
// endpoint from CLI flags.
func BuildDeleteUserPayload(sSAServerDeleteUserBody string, sSAServerDeleteUserUserID string) (*ssaserver.DeleteUserPayload, error) {
	var err error
	var body DeleteUserRequestBody
	{
		err = json.Unmarshal([]byte(sSAServerDeleteUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"password\": \"pass12345\"\n   }'")
		}
	}
	var userID int
	{
		var v int64
		v, err = strconv.ParseInt(sSAServerDeleteUserUserID, 10, 64)
		userID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for userID, must be INT")
		}
	}
	v := &ssaserver.DeleteUserPayload{
		Password: body.Password,
	}
	v.UserID = userID
	return v, nil
}

// BuildSaveDataPayload builds the payload for the SSAServer Save_data endpoint
// from CLI flags.
func BuildSaveDataPayload(sSAServerSaveDataBody string, sSAServerSaveDataGroupID string) (*ssaserver.SaveDataPayload, error) {
	var err error
	var body SaveDataRequestBody
	{
		err = json.Unmarshal([]byte(sSAServerSaveDataBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"Data\": \"UmVydW0gcmVwcmVoZW5kZXJpdCBxdWkgbW9sZXN0aWFlLg==\",\n      \"Image\": \"UXVpcyBvcHRpbyByZXJ1bSBjb3JydXB0aSBtb2xsaXRpYSB1dCBkb2xvcmUu\",\n      \"data_name\": \"Diary_312_2019-03-02_12-07-35\",\n      \"data_type\": 1,\n      \"image_name\": \"Image_2017-05-25-26-32\",\n      \"password\": \"pass12345\",\n      \"title\": \"たいとる\",\n      \"user_id\": 28532\n   }'")
		}
		if body.Data == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("Data", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var groupID string
	{
		groupID = sSAServerSaveDataGroupID
	}
	v := &ssaserver.SaveDataPayload{
		UserID:    body.UserID,
		Password:  body.Password,
		DataName:  body.DataName,
		DataType:  body.DataType,
		Data:      body.Data,
		Title:     body.Title,
		ImageName: body.ImageName,
		Image:     body.Image,
	}
	v.GroupID = groupID
	return v, nil
}

// BuildReturnDataListPayload builds the payload for the SSAServer
// Return_data_list endpoint from CLI flags.
func BuildReturnDataListPayload(sSAServerReturnDataListBody string, sSAServerReturnDataListGroupID string) (*ssaserver.ReturnDataListPayload, error) {
	var err error
	var body ReturnDataListRequestBody
	{
		err = json.Unmarshal([]byte(sSAServerReturnDataListBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"user_id\": 537829\n   }'")
		}
	}
	var groupID string
	{
		groupID = sSAServerReturnDataListGroupID
	}
	v := &ssaserver.ReturnDataListPayload{
		UserID: body.UserID,
	}
	v.GroupID = groupID
	return v, nil
}

// BuildPickUpDataPayload builds the payload for the SSAServer Pick_up_data
// endpoint from CLI flags.
func BuildPickUpDataPayload(sSAServerPickUpDataBody string, sSAServerPickUpDataGroupID string, sSAServerPickUpDataDataUserID string) (*ssaserver.PickUpDataPayload, error) {
	var err error
	var body PickUpDataRequestBody
	{
		err = json.Unmarshal([]byte(sSAServerPickUpDataBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"data_name\": \"Record_12_2019-06-02_12-07-35\",\n      \"data_type\": 0,\n      \"image_name\": \"Image_2017-05-25-26-32\",\n      \"user_id\": 65\n   }'")
		}
	}
	var groupID string
	{
		groupID = sSAServerPickUpDataGroupID
	}
	var dataUserID int
	{
		var v int64
		v, err = strconv.ParseInt(sSAServerPickUpDataDataUserID, 10, 64)
		dataUserID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for dataUserID, must be INT")
		}
	}
	v := &ssaserver.PickUpDataPayload{
		DataType:  body.DataType,
		UserID:    body.UserID,
		DataName:  body.DataName,
		ImageName: body.ImageName,
	}
	v.GroupID = groupID
	v.DataUserID = dataUserID
	return v, nil
}
