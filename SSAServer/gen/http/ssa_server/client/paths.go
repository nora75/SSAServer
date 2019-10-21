// Code generated by goa v3.0.3, DO NOT EDIT.
//
// HTTP request path constructors for the SSAServer service.
//
// Command:
// $ goa gen SSAServer/design

package client

import (
	"fmt"
)

// RegisterSSAServerPath returns the URL path to the SSAServer service Register HTTP endpoint.
func RegisterSSAServerPath() string {
	return "/Registration"
}

// LoginSSAServerPath returns the URL path to the SSAServer service Login HTTP endpoint.
func LoginSSAServerPath() string {
	return "/Login"
}

// ChangeGroupSSAServerPath returns the URL path to the SSAServer service Change_group HTTP endpoint.
func ChangeGroupSSAServerPath(userID int) string {
	return fmt.Sprintf("/users/%v", userID)
}

// DeleteUserSSAServerPath returns the URL path to the SSAServer service Delete_user HTTP endpoint.
func DeleteUserSSAServerPath(userID int) string {
	return fmt.Sprintf("/users/%v", userID)
}

// SaveDataSSAServerPath returns the URL path to the SSAServer service Save_data HTTP endpoint.
func SaveDataSSAServerPath(groupID string) string {
	return fmt.Sprintf("/group/%v", groupID)
}

// ReturnDataListSSAServerPath returns the URL path to the SSAServer service Return_data_list HTTP endpoint.
func ReturnDataListSSAServerPath(groupID string) string {
	return fmt.Sprintf("/group/%v", groupID)
}

// PickUpDataSSAServerPath returns the URL path to the SSAServer service Pick_up_data HTTP endpoint.
func PickUpDataSSAServerPath(groupID string, dataUserID int) string {
	return fmt.Sprintf("/group/%v/%v", groupID, dataUserID)
}