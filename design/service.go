package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("SSAServer", func() {
	Description("Service for account management.")
	Error("Invalid ID")
	Error("Invalid Password")
	Error("Invalid Group ID")
	Error("Invalid Request")
	Error("Invalid User Name")
	Error("Invalid data_name")
	Error("Invalid Data")
	Error("Invalid mail-address")
	Error("Require Authentication")
	Error("The user already exists")

	Method("Register", func() {
		Payload(func() {
			Attribute("user_name", String, "User Name")
			Attribute("password", String, "User Password")
			Attribute("mail", String, "User mail-address")
			Attribute("group_id", String, "Gourp ID")
			Required("user_name", "password", "mail")
			Example("Register example without group_id", func(){
				Description("The sample of register")
				Value(map[string]string{"user_name": "hoge",
					"password": "password",
					"mail": "hoge@hoge.com"})
			})
			Example("Register example with group_id", func(){
				Description("The sample of register")
				Value(map[string]string{"user_name": "hoge",
					"password": "password",
					"mail": "hoge@hoge.com",
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
			POST("/Registration")
			Response(StatusOK)
			Response("Invalid Group ID", StatusNotFound)
			Response("The user already exists", StatusBadRequest)
			Response("Invalid Request", StatusBadRequest)
		})

	})

	Method("Login", func() {
		Payload(func() {
			Username("mail", String)
			Password("password", String)
			Attribute("group_id", Int, "Gourp ID")
			Attribute("id", Int, "User ID")
			Required("mail", "password")
			Example("Login example without group_id", func(){
				Description("The sample of login")
				Value(map[string]string{"mail": "hoge@hoge.com",
					"password": "password"})
			})
			Example("Login example with group_id", func(){
				Description("The sample of login")
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
			Response("Invalid Request", StatusBadRequest)
			Response("Invalid Group ID", StatusNotFound)
		})

	})

	Method("Join_group", func(){
		Payload(func(){
			Username("id", Int, "User ID")
			Password("password", String, "User password")
			Attribute("group_id", String, "Group ID")
			Required("id", "password","group_id")
			Example("Join group1", func(){
				Description("Join group1")
				Value(map[string]string{"id": "123",
					"password": "password",
					"group_id": "group1"})
			})
			Example("Join group2", func(){
				Description("Join group1")
				Value(map[string]string{"id": "3123",
					"password": "password",
					"group_id": "group2"})
			})
		})

		HTTP(func(){
			POST("/{id}")
			Response("Invalid Group ID", StatusBadRequest)
			Response("Invalid Request", StatusBadRequest)
			Response(StatusOK)
		})
	})

	Method("Delete_user", func(){
		Payload(func(){
			Username("id", Int)
			Password("password", String)
			Required("id", "password")
			Example("Delete user 3123", func(){
				Description("Delete user who's id is 3123")
				Value(map[string]string{"password": "password"})
			})
		})

		HTTP(func(){
			DELETE("/{id}")
			Response("Invalid Request", StatusBadRequest)
			Response(StatusOK)
		})
	})

	Method("Save", func() {
		Payload(func() {
			Attribute("group_id", String, "Gourp ID")
			Attribute("id", Int, "User ID")
			Attribute("data_name", String, "Data name")
			Attribute("data_type", Int, "Data name")
			Attribute("Data1", Any, "Data1")
			Attribute("title", String, "Diary title")
			Attribute("image_name", String, "Image name")
			Attribute("Data2", Any, "Data2")
			Required("group_id", "id", "data_name", "Data1")
		})

		HTTP(func() {
			POST("/{group_id}")
			Response("Invalid Group ID", StatusNotFound)
			Response("Invalid Request", StatusBadRequest)
			Response("Invalid Data", StatusBadRequest)
			Response(StatusOK)
		})

	})

	Method("Return_data_list", func() {
		Payload(func() {
			Attribute("group_id", String, "Gourp ID")
			Attribute("id", Int, "User ID")
			Required("group_id", "id")
		})

		Result(MyResultType, func(){
			View("data_list")
		})

		HTTP(func() {
			GET("/{group_id}")
			Response(StatusOK)
			Response("Invalid Group ID", StatusNotFound)
			Response("Invalid Request", StatusBadRequest)
		})

	})

	Method("Pick_up", func() {
		Payload(func() {
			Attribute("group_id", String, "Gourp ID")
			Attribute("id", Int, "User ID")
			Attribute("data_name", String, "Data name")
			Attribute("data_type", String, "Data type")
			Required("group_id", "id", "data_name")
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
			GET("/{group_id}")
			Response(StatusOK)
			Response("Invalid Group ID", StatusNotFound)
			Response("Invalid Request", StatusBadRequest)
			Response("Invalid data_name", StatusBadRequest)
		})

	})

	Files("/openapi.json", "../../gen/http/openapi.json")
})
