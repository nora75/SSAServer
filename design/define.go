package design

import (
	. "goa.design/goa/v3/dsl"
)

// BasicAuth defines a security scheme using basic authentication. The scheme
// protects the "signin" action used to create JWTs.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during signin")
	Scope("api:read", "Read-only access")
})

// MyResultType is declation of use in result views
var MyResultType = ResultType("application/vnd.goa.ssa", func(){
	ContentType("application/json")

	Attributes(func() {
		Attribute("user_id", String, "User Name", func(){
			Example("The first user id", func(){
				Description("The user who has register first")
				Value(Val{"id": 1})
			})
			Example("The 234's user id", func(){
				Description("The user who has register 234's")
				Value(Val{"id": 234})
			})
		})
		Attribute("user_name", String, "User Name", func(){
			Example("The user's name Hirano", func(){
				Description("Simple usecase")
				Value(Val{"user_name": "Hirano"})
			})
			Example("The user's name Hirano Tatsuya", func(){
				Description("Can contain space")
				Value(Val{"user_name": "Hirano Tatsuya"})
			})
			Example("The user's name 平野 竜也", func(){
				Description("Can use utf-8")
				Value(Val{"user_name": "平野 竜也"})
			})
		})
		Attribute("password", String, "User Password", func(){
			Example("The password is 'password'", func(){
				Description("The password is 'password' but shoudn't use it")
				Value(Val{"password": "password"})
			})
			Example("The password is 'PassW0rd!!@'", func(){
				Description("The password is 'PassW0rd!!@'")
				Value(Val{"password": "PassW0rd!!@"})
			})
		})
		Attribute("mail", String, "User mail-address", func(){
			Example("Example of mail-address", func(){
				Description("Example of mail-address")
				Value(Val{"mail": "hoge@hoge.com"})
			})
		})
		Attribute("group_id", String, "Gourp ID", func(){
			Example("The group name is 'bzuiphjjgdas'", func(){
				Description("The user who belong to Group 'bzuiphjjgdas'")
				Value(Val{"gourp_id": "bzuiphjjgdas"})
			})
		})
		Attribute("data_name", String, "Data name", func(){
			Example("The example of record data", func(){
				Description("The record data that's name is 'Record_12_2019-06-02_12-07-35'")
				Value(Val{"data_name": "Record_12_2019-06-02_12-07-35"})
			})
			Example("The example of Diary data", func(){
				Description("The diary data that's name is 'diary_312_2019-03-02_12-07-35'")
				Value(Val{"data_name": "Record_312_2019-03-02_12-07-35"})
			})
		})
		Attribute("Data", Any, "Data")
		Attribute("data_type", String, "Data's name", func(){
			Example("The record", func(){
				Description("The data_type 0(record)")
				Value(Val{"data_type": 0})
			})
			Example("The Diary", func(){
				Description("The data_type 1(diary)")
				Value(Val{"data_type": 1})
			})
		})
	})

	View("default", func() {
		Attribute("mail")
		Attribute("user_name")
	})

	View("extended", func() {
		Attribute("mail")
		Attribute("user_name")
		Attribute("group_id")
	})

	View("login", func() {
		Attribute("user_id")
		Attribute("mail")
		Attribute("user_name")
	})

	View("extended", func() {
		Attribute("user_id")
		Attribute("mail")
		Attribute("user_name")
		Attribute("group_id")
	})

	View("group_id", func() {
		Attribute("group_id")
	})
	
	View("data", func() {
		Attribute("Data")
		Attribute("data_type")
	})
})

var _ = API("SSA", func() {
	Title("SSAServer")
	Description("Service for record talking.")
	Version("1.0")
	Server("SSAServer", func() {
		Host("localhost", func() {
			URI("http://localhost:8000/")
		})
		Services("SSAServer")
	})
})

var _ = Service("SSAServer", func() {
	Description("Service for account management.")
	Error("Invalid ID")
	Error("Invalid Password")
	Error("Invalid Group ID")
	Error("Invalid Request")
	Error("Invalid User Name")
	Error("Invalid data_name")
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
				Value(Val{"user_name": "hoge",
					"password": "password",
					"mail": "hoge@hoge.com"})
			})
			Example("Register example with group_id", func(){
				Description("The sample of register")
				Value(Val{"user_name": "hoge",
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
			Required("mail", "password")
			Example("Login example without group_id", func(){
				Description("The sample of login")
				Value(Val{"mail": "hoge@hoge.com",
					"password": "password"})
			})
			Example("Login example with group_id", func(){
				Description("The sample of login")
				Value(Val{"mail": "hoge@hoge.com",
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
		Payload(func() {
			Security(BasicAuth)
			Username("id", Int, "User ID")
			Password("password", Int, "User password")
			Required("id", "mail")
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
				Value(Val{"id": "123",
					"password": "password",
					"group_id": "group1"})
			})
			Example("Join group2", func(){
				Description("Join group1")
				Value(Val{"id": "3123",
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
		Payload(func(){
		Security(BasicAuth)
			Username("id", Int)
			Password("password", String)
			Required("id", "password")
			Example("Delete user 3123", func(){
				Description("Delete user who's id is 3123")
				Value(Val{"password": "password"})
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
			Attribute("data_type", String, "Data name")
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
