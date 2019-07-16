package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("SSAServer", func() {
	Description("SSAアプリケーションのサーバーサイド")
	Error("Invalid_User_ID", CustomErrorType)
	Error("Invalid_Password", CustomErrorType)
	Error("Invalid_Group_ID", CustomErrorType)
	Error("Invalid_Request", CustomErrorType)
	Error("Invalid_User_Name", CustomErrorType)
	Error("Invalid_data_name", CustomErrorType)
	Error("Invalid_Data", CustomErrorType)
	Error("Invalid_mail-address", CustomErrorType)
	Error("Require_Authentication", CustomErrorType)
	Error("The_user_already_exists", CustomErrorType)

	Method("Register", func() {
		Description("SSAへの新規登録")
		Payload(func() {
			Attribute("user_name", String, "User Name", func(){
				Example("平野竜也")
			})
			Attribute("password", String, "User Password", func(){
				Example("passW0rd")
			})
			Attribute("mail", String, "User mail-address", func(){
				Format("email")
				Example("hoge@hoge.com")
			})
			Attribute("group_id", String, "Group ID", func(){
				Example("group-bzuiphjjgdas")
			})
			Required("user_name", "password", "mail")
		})

		Result(MyResultType, func(){
			View("extended")
			Required("user_id","mail")
		})

		HTTP(func() {
			POST("/Registration")
			Response(StatusOK)
			Response("Invalid_Group_ID", StatusNotFound)
			Response("The_user_already_exists", StatusBadRequest)
			Response("Invalid_Request", StatusBadRequest)
		})

	})

	Method("Login", func() {
		Description("SSAへのログイン")
		Payload(func() {
			Attribute("mail", String, "User mail-address", func(){
				Example("hoge@hoge.com")
			})
			Attribute("password", String, "User Password", func(){
				Example("password")
			})
			Required("mail", "password")
		})

		Result(Boolean)

		HTTP(func() {
			POST("/Login")
			Response(StatusOK)
			Response("Invalid_Request", StatusBadRequest)
			Response("Invalid_Group_ID", StatusNotFound)
		})

	})

	Method("Change_group", func(){
		Description("グループIDを変更する")
		Payload(func(){
			Attribute("user_id", Int, "User ID", func(){
				Example(123)
			})
			Attribute("password", String, "User password", func(){
				Example("pass1234")
			})
			Attribute("group_id", String, "Group ID", func(){
				Example("group2")
			})
			Required("user_id", "password","group_id")
		})

		Result(Boolean)

		HTTP(func(){
			POST("/users/{user_id}")
			Response("Invalid_Group_ID", StatusBadRequest)
			Response("Invalid_Request", StatusBadRequest)
			Response(StatusOK)
		})
	})

	Method("Delete_user", func(){
		Description("既存ユーザーの消去")
		Payload(func(){
			Attribute("user_id", Int, "User ID", func(){
				Example(1342)
			})
			Attribute("password", String, "User Password", func(){
				Example("pass12345")
			})
			Required("user_id", "password")
		})

		Result(Boolean)

		HTTP(func(){
			DELETE("/users/{user_id}")
			Response("Invalid_Request", StatusBadRequest)
			Response(StatusOK)
		})
	})

	Method("Save_data", func() {
		Description("データをサーバーへ保存する")
		Payload(func() {
			Attribute("group_id", String, "Group ID", func(){
				Example("group1")
			})
			Attribute("user_id", Int, "User ID", func(){
				Example(28532)
			})
			Attribute("data_name", String, "Data name", func(){
				Example("Diary_312_2019-03-02_12-07-35")
			})
			Attribute("data_type", Int, "Data name", func(){
				Example(1)
			})
			Attribute("Data", Any, "Data", func(){
				Meta("swagger:example", "false")
			})
			Attribute("title", String, "Diary title", func(){
				Example("たいとる")
			})
			Attribute("image_name", String, "Image name", func(){
				Example("Image_2017-05-25-26-32")
			})
			Attribute("Image", Any, "Image", func(){
				Meta("swagger:example", "false")
			})
			Required("group_id", "user_id", "data_name", "data_type", "Data")
		})

		Result(Boolean)

		HTTP(func() {
			POST("/group/{group_id}")
			MultipartRequest()
			Response("Invalid_Group_ID", StatusNotFound)
			Response("Invalid_Request", StatusBadRequest)
			Response("Invalid_Data", StatusBadRequest)
			Response(StatusOK)
		})

	})

	Method("Return_data_list", func() {
		Description("データのリストを取得する")
		Payload(func() {
			Attribute("group_id", String, "Group ID", func(){
				Example("group-agieoa")
			})
			Attribute("user_id", Int, "User ID", func(){
				Example(537829)
			})
			Required("group_id", "user_id")
		})

		Result(DataResult)

		HTTP(func() {
			GET("/group/{group_id}")
			Response(StatusOK)
			Response("Invalid_Group_ID", StatusNotFound)
			Response("Invalid_Request", StatusBadRequest)
		})

	})

	Method("Pick_up_data", func() {
		Description("データをサーバーから取得する")
		Payload(func() {
			Attribute("group_id", String, "Group ID", func(){
				Example("group-isg")
			})
			Attribute("data_type", String, "Data type", func(){
				Example("0")
			})
			Attribute("user_id", Int, "User ID", func(){
				Example(65)
			})
			Attribute("data_name", String, "Data name", func(){
				Example("Record_12_2019-06-02_12-07-35")
			})
			Required("group_id", "user_id", "data_name")
		})

		Result(MyResultType, func(){
			View("data_extended_with_image")
			Required("Data", "data_type", "data_name")
		})

		HTTP(func() {
			GET("/group/{group_id}/{data_type}")
			MultipartRequest()
			Response(StatusOK)
			Response("Invalid_Group_ID", StatusNotFound)
			Response("Invalid_Request", StatusBadRequest)
			Response("Invalid_data_name", StatusBadRequest)
		})

	})

	Files("/openapi.json", "gen/http/openapi.json")
})
