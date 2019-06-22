package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("SSAServer", func() {
	Description("SSAアプリケーションのサーバーサイド")
	Error("Invalid_User_ID")
	Error("Invalid_Password")
	Error("Invalid_Group_ID")
	Error("Invalid_Request")
	Error("Invalid_User_Name")
	Error("Invalid_data_name")
	Error("Invalid_Data")
	Error("Invalid_mail-address")
	Error("Require_Authentication")
	Error("The_user_already_exists")

	Method("Register", func() {
		Description("SSAへの新規登録")
		Payload(func() {
			Attribute("user_name", String, "User Name", func(){
				Example("The user's name Hirano Tatsuya", func(){
					Description("単純な例")
					Value("Hirano Tatsuya")
				})
				Example("The user's name 平野 竜也", func(){
					Description("UTF-8に対応")
					Value("平野 竜也")
				})
			})
			Attribute("password", String, "User Password", func(){
				Example("The password is 'password'", func(){
					Description("'password'がパスワードの場合")
					Value("password")
				})
				Example("The password is 'PassW0rd!!@'", func(){
					Description("'PassW0rd!!@'がパスワードの場合")
					Value("PassW0rd!!@")
				})
			})
			Attribute("mail", String, "User mail-address", func(){
				Example("Example of mail-address", func(){
					Description("メアド")
					Value("hoge@hoge.com")
				})
			})
			Attribute("group_id", String, "Gourp ID", func(){
				Example("'bzuiphjjgdas'がグループ名の場合", func(){
					Value("bzuiphjjgdas")
				})
			})
			Required("user_name", "password", "mail")
			Example("group_id抜きの例", func(){
				Value(map[string]string{
					"user_name": "hoge",
					"password": "password",
					"mail": "hoge@hoge.com",
				})
			})
			Example("group_id付きの例", func(){
				Value(map[string]string{
					"user_name": "hoge",
					"password": "password",
					"mail": "hoge@hoge.com",
					"group_id": "group1",
				})
			})
		})

		Result(MyResultType, func(){
			View("default")
		})

		Result(MyResultType, func(){
			View("extended")
			Required("group_id")
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
			Attribute("mail", MapOf(String, String), "User mail-address")
			Attribute("password", MapOf(String, String), "User Password")
			Attribute("group_id", MapOf(String, Int), "Gourp ID")
			Attribute("user_id", MapOf(String, Int), "User ID")
			Required("mail", "password")
			Example("group_id抜きの場合", func(){
				Value(map[string]string{"mail": "hoge@hoge.com",
					"password": "password"})
			})
			Example("group_id付きの場合", func(){
				Value(map[string]string{"mail": "hoge@hoge.com",
					"password": "password",
					"group_id": "group1"})
			})
		})

		Result(MyResultType, func(){
			View("default")
		})

		Result(MyResultType, func(){
			View("extended")
			Required("group_id")
		})

		HTTP(func() {
			POST("/Login")
			Response(StatusOK)
			Response("Invalid_Request", StatusBadRequest)
			Response("Invalid_Group_ID", StatusNotFound)
		})

	})

	Method("Change_group", func(){
		Description("グループIDを変更する場合")
		Payload(func(){
			Attribute("user_id", Int, "User ID")
			Attribute("password", MapOf(String, String), "User password")
			Attribute("group_id", MapOf(String, String), "Group ID")
			Required("user_id", "password","group_id")
			Example("'group2'へ変更", func(){
				Value(map[string]string{
					"password": "password",
					"group_id": "group2",
				})
			})
		})

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
			Attribute("user_id", Int, "User ID")
			Attribute("password", MapOf(String, String), "User Password")
			Required("user_id", "password")
			Example("ユーザー:'3123'の消去", func(){
				Value(map[string]string{"password": "password"})
			})
		})

		HTTP(func(){
			DELETE("/users/{user_id}")
			Response("Invalid_Request", StatusBadRequest)
			Response(StatusOK)
		})
	})

	Method("Save", func() {
		Description("データをサーバーへ保存する")
		Payload(func() {
			Attribute("group_id", String, "Gourp ID")
			Attribute("user_id", MapOf(String, Int), "User ID")
			Attribute("data_name", MapOf(String, String), "Data name")
			Attribute("data_type", MapOf(String, Int), "Data name")
			Attribute("Data1", MapOf(String, Any), "Data1")
			Attribute("title", MapOf(String, String), "Diary title")
			Attribute("image_name", MapOf(String, String), "Image name")
			Attribute("Data2", MapOf(String, Any), "Data2")
			Required("group_id", "user_id", "data_name", "Data1")
		})

		HTTP(func() {
			POST("/group/{group_id}")
			Response("Invalid_Group_ID", StatusNotFound)
			Response("Invalid_Request", StatusBadRequest)
			Response("Invalid_Data", StatusBadRequest)
			Response(StatusOK)
		})

	})

	Method("Return_data_list", func() {
		Description("データのリストを取得する")
		Payload(func() {
			Attribute("group_id", String, "Gourp ID")
			Attribute("user_id", MapOf(String, Int), "User user_ID")
			Required("group_id", "user_id")
		})

		Result(MyResultType, func(){
			View("data_list")
		})

		HTTP(func() {
			GET("/group/{group_id}")
			Response(StatusOK)
			Response("Invalid_Group_ID", StatusNotFound)
			Response("Invalid_Request", StatusBadRequest)
		})

	})

	Method("Pick_up", func() {
		Description("データをサーバーから取得する")
		Payload(func() {
			Attribute("group_id", String, "Gourp ID")
			Attribute("user_id", MapOf(String, Int), "User ID")
			Attribute("data_name", MapOf(String, String), "Data name")
			Attribute("data_type", MapOf(String, String), "Data type")
			Required("group_id", "user_id", "data_name")
		})

		Result(MyResultType, func(){
			View("data")
		})

		Result(MyResultType, func(){
			View("data_extended")
			Required("title")
		})

		Result(MyResultType, func(){
			View("data_extended_with_image")
			Required("title", "image_name")
		})

		HTTP(func() {
			GET("/group/{group_id}")
			Response(StatusOK)
			Response("Invalid_Group_ID", StatusNotFound)
			Response("Invalid_Request", StatusBadRequest)
			Response("Invalid_data_name", StatusBadRequest)
		})

	})

	Files("/openapi.json", "../../gen/http/openapi.json")
})
