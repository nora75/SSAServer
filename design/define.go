package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("SSA", func() {
	Title("SSAServer")
	Description("Service for record talking.")
	Version("1.0")
	Server("SSAServer", func() {
		Host("localhost", func() {
			URI("http://localhost:8000/")
		})
		HTTP(func(){
			Path("SSAServer")
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

		Result(Default, func(){
			View("default")
		})

		Result(Extended, func(){
			View("extended")
			Required("group_id")
		})

		HTTP(func() {
			POST("/Register")
			Response(StatusOK)
			Response("Bad Group ID", NotFound)
			Response("The user already exists", BadRequest)
		})

	})

	Method("Login", func() {
		Security(Basic)
		Payload(func() {
			Username("mail", String)
			Password("password", String)
			Attribute("group_id", Int, "Gourp ID")
			Required("mail", "password")
		})

		Result(Default, func(){
			View("default")
		})

		Result(Extended, func(){
			View("extended")
			Required("group_id")
		})

		HTTP(func() {
			GET("/Login")
			Response(StatusOK)
			Response("Bad Group ID", NotFound)
			Response("Bad Password", BadRequest)
			Response("Bad User Name", BadRequest)
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
			Response("Bad Group ID", BadRequest)
			Response(StatusOK)
		})
	})

	Method("Save", func() {
		Payload(func() {
			Attribute("id", Int, "User ID")
			Attribute("group_id", String, "Gourp ID")
			Attribute("data_name", String, "Data name")
			Attribute("Data", String, "Data name")
			Required("id", "group_id", "data_name", "Data")
		})

		Result(Boolean)

		HTTP(func() {
			GET("/Save/{id}")
			Response("Bad Password", BadRequest)
			Response("Bad Request", BadRequest)
			Response("Bad User Name", BadRequest)
			Response("Bad Group ID", NotFound)
			Response(StatusOK)
		})

	})

	Method("PickUp", func() {
		Payload(func() {
			Attribute("id", Int, "User ID")
			Attribute("group_id", Strign, "Gourp ID")
			Attribute("data_name", String, "Data name")
			Required("id", "group_id", "data_name")
		})

		Result(String)

		HTTP(func() {
			GET("/PickUp/{1}")
			Response(StatusOK)
			Response("Bad Group ID", NotFound)
			Response("Bad Request", BadRequest)
			Response("Bad data_name", BadRequest)
		})

	})

	Method("Generate", func() {
		Payload(func() {
			Attribute("id", Int, "User ID")
			Attribute("mail", Int, "User mail-address")
			Required("id", "mail")
		})

		Result(GroupID, func(){
			View("group_id")
		})

		HTTP(func() {
			GET("/GroupID/{id}")
			Response(StatusOK)
			Response("Bad ID", BadRequest)
			Response("Bad mail-address", BadRequest)
			Response("Bad Request", Forbidden)
		})

	})

	Files("/openapi.json", "../../gen/http/openapi.json")
})
