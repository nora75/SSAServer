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
			URI("http://localhost:8000/SSAServer")
		})
		// List the services hosted by this server.
		Services("SSAServer")
	})
})

var _ = Service("SSAServer", func() {
	Description("Service for account management.")

	Method("Register", func() {
		Payload(func() {
			Field(1, "user_name", String, "User Name")
			Field(2, "password", String, "User Password")
			Field(3, "mail", String, "User mail-address")
			Field(4, "group_id", Int, "Gourp ID")
			Required("user_name", "password", "mail")
		})

		Result(Boolean)

		HTTP(func() {
			GET("/Register")
            Response(StatusOK)
            Error("The user already exists", BadRequest)
            Error("Bad Group ID", BadRequest)
		})

	})

	Method("Login", func() {
		Security(Basic)
		Payload(func() {
			Username("mail", String)
			Password("password", String)
			Field(3, "group_id", Int, "Gourp ID")
			Required("mail", "password")
		})

		Result(Boolean)

		HTTP(func() {
			GET("/Login")
            Response(StatusOK)
            Error("Bad User Name", BadRequest)
            Error("Bad Password", BadRequest)
            Error("Bad Group ID", BadRequest)
		})

	})

	Method("Join", func(){
		Payload(func(){
			Field(1,"group_id", Int, "Group ID")
		})

		Result(Int)

		HTTP(func(){
			GET("/Join")
            Response(StatusOK)
            Error("Bad Group ID", BadRequest)
		})
	})

	Method("Save", func() {
		Payload(func() {
			Field(1, "id", Int, "User ID")
			Field(2, "group_id", Int, "Gourp ID")
			Field(3, "data_name", String, "Data name")
			Field(4, "Data", String, "Data name")
			Required("id", "group_id", "data_name", "Data")
		})

		Result(Boolean)

		HTTP(func() {
			GET("/Save")
            Response(StatusOK)
            Error("Bad User Name", BadRequest)
            Error("Bad Password", BadRequest)
		})

	})

	Method("PickUp", func() {
		Payload(func() {
			Field(1, "id", Int, "User ID")
			Field(2, "group_id", Int, "Gourp ID")
			Field(3, "data_name", String, "Data name")
			Required("id", "group_id", "data_name")
		})

		Result(String)

		HTTP(func() {
			GET("/PickUp")
            Response(StatusOK)
            Error("Bad Group ID", BadRequest)
            Error("Bad data_name", BadRequest)
		})

	})

	Method("Generate", func() {
		Payload(func() {
			Field(1, "id", Int, "User ID")
			Field(2, "mail", Int, "User mail-address")
			Required("id", "mail")
		})

		Result(Int)

		HTTP(func() {
			GET("/GroupID")
            Response(StatusOK)
            Error("Bad ID", BadRequest)
            Error("Bad mail-address", BadRequest)
		})

	})

	Files("/openapi.json", "../../gen/http/openapi.json")
})
