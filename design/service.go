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

		Result(Boolean)

		HTTP(func() {
			GET("/Register")
			Response(StatusOK)
			Response("Invalid Group ID", StatusNotFound)
			Response("The user already exists", StatusBadRequest)
			Response("Invalid Request", StatusBadRequest)
		})

	})

	Method("Login", func() {
		Security(BasicAuth)
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
			View("login")
		})

		Result(MyResultType, func(){
			View("login_extended")
			Required("group_id")
		})

		HTTP(func() {
			POST("/{id}")
			Response(StatusOK)
			Response("Invalid Request", StatusBadRequest)
			Response("Invalid Group ID", StatusNotFound)
		})

	})

	Method("Generate_group_id", func() {
		Security(BasicAuth)
		Payload(func() {
			Username("id", Int, "User ID")
			Password("password", Int, "User password")
			Required("id", "password")
		})

		Result(MyResultType, func(){
			View("group_id")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("Invalid ID", StatusBadRequest)
			Response("Invalid Request", StatusBadRequest)
		})

	})

	Method("Join_group", func(){
		Payload(func(){
			Attribute("id", Int, "User ID")
			Attribute("mail", String, "User mail-address")
			Attribute("group_id", String, "Group ID")
			Required("id", "mail","group_id")
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

		Result(Boolean)

		HTTP(func(){
			POST("/{id}/Join")
			Response("Invalid Group ID", StatusBadRequest)
			Response("Invalid Request", StatusBadRequest)
			Response(StatusOK)
		})
	})

	Method("Delete_user", func(){
		Security(BasicAuth)
		Payload(func(){
			Username("id", Int)
			Password("password", String)
			Required("id", "password")
			Example("Delete user 3123", func(){
				Description("Delete user who's id is 3123")
				Value(map[string]string{"password": "password"})
			})
		})

		Result(Boolean)

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
			Attribute("Data", Any, "Data name")
			Required("group_id", "id", "data_name", "Data")
		})

		Result(Boolean)

		HTTP(func() {
			POST("/{group_id}")
			Response("Invalid Group ID", StatusNotFound)
			Response("Invalid Request", StatusBadRequest)
			Response("Invalid Data", StatusBadRequest)
			Response(StatusOK)
		})

	})

	Method("Pick_up", func() {
		Payload(func() {
			Attribute("group_id", String, "Gourp ID")
			Attribute("id", Int, "User ID")
			Attribute("data_name", String, "Data name")
			Required("group_id", "id", "data_name")
		})

		Result(MyResultType, func(){
			View("data")
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
