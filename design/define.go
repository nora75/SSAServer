package design

import (
	. "goa.design/goa/v3/dsl"
)

var BasicAuth = BasicAuthSecurity("basic_auth")

var MyResultType = ResultType("application/vnd.goa.ssa", func(){
	ContentType("application/json")

	Attributes(func() {
		Attribute("user_name", String, "User Name")
		Attribute("password", String, "User Password")
		Attribute("mail", String, "User mail-address")
		Attribute("group_id", String, "Gourp ID")
		Attribute("data_name", String, "Data name")
		Attribute("Data", Any, "Data name")
	})

	View("default", func() {
		// "id" and "name" must be result type attributes
		Attribute("mail")
		Attribute("user_name")
	})

	View("extended", func() {
		Attribute("mail")
		Attribute("user_name")
		Attribute("group_id")
	})

	View("group_id", func() {
		Attribute("group_id")
	})
	
	View("data", func() {
		Attribute("Data")
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
	Error("The user already exists")

	Method("Register", func() {
		Payload(func() {
			Attribute("user_name", String, "User Name")
			Attribute("password", String, "User Password")
			Attribute("mail", String, "User mail-address")
			Attribute("group_id", String, "Gourp ID")
			Required("user_name", "password", "mail")
		})

		Result(MyResultType, func(){
			View("default")
		})

		Result(MyResultType, func(){
			View("extended")
			Required("group_id")
		})

		HTTP(func() {
			POST("/Register")
			Response(StatusOK)
			Response("Invalid Group ID", StatusNotFound)
			Response("The user already exists", StatusBadRequest)
		})

	})

	Method("Login", func() {
		Security(BasicAuth)
		Payload(func() {
			Username("mail", String)
			Password("password", String)
			Attribute("group_id", Int, "Gourp ID")
			Required("mail", "password")
		})

		Result(MyResultType, func(){
			View("default")
		})

		Result(MyResultType, func(){
			View("extended")
			Required("group_id")
		})

		HTTP(func() {
			GET("/Login")
			Response(StatusOK)
			Response("Invalid Group ID", StatusNotFound)
			Response("Invalid Password", StatusBadRequest)
			Response("Invalid User Name", StatusBadRequest)
		})

	})

	Method("Join", func(){
		Payload(func(){
			Attribute("id", Int, "User ID")
			Attribute("group_id", String, "Group ID")
		})

		Result(Boolean)

		HTTP(func(){
			GET("/Join")
			Response("Invalid Group ID", StatusBadRequest)
			Response(StatusOK)
		})
	})

	Method("Save", func() {
		Payload(func() {
			Attribute("id", Int, "User ID")
			Attribute("group_id", String, "Gourp ID")
			Attribute("data_name", String, "Data name")
			Attribute("Data", Any, "Data name")
			Required("id", "group_id", "data_name", "Data")
		})

		Result(Boolean)

		HTTP(func() {
			GET("/Save/{id}")
			Response("Invalid Password", StatusBadRequest)
			Response("Invalid Request", StatusBadRequest)
			Response("Invalid User Name", StatusBadRequest)
			Response("Invalid Group ID", StatusNotFound)
			Response(StatusOK)
		})

	})

	Method("PickUp", func() {
		Payload(func() {
			Attribute("id", Int, "User ID")
			Attribute("group_id", String, "Gourp ID")
			Attribute("data_name", String, "Data name")
			Required("id", "group_id", "data_name")
		})

		Result(MyResultType, func(){
			View("data")
		})

		HTTP(func() {
			GET("/PickUp")
			Response(StatusOK)
			Response("Invalid Group ID", StatusNotFound)
			Response("Invalid Request", StatusBadRequest)
			Response("Invalid data_name", StatusBadRequest)
		})

	})

	Method("Generate", func() {
		Payload(func() {
			Attribute("id", Int, "User ID")
			Attribute("mail", Int, "User mail-address")
			Required("id", "mail")
		})

		Result(MyResultType, func(){
			View("group_id")
		})

		HTTP(func() {
			GET("/GroupID/{id}")
			Response(StatusOK)
			Response("Invalid ID", StatusBadRequest)
			Response("Invalid mail-address", StatusBadRequest)
			Response("Invalid Request", StatusForbidden)
		})

	})

	Files("/openapi.json", "../../gen/http/openapi.json")
})
