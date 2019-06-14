package design

import (
	. "goa.design/goa/v3/dsl"
)

var MyResultType = ResultType("application/vnd.goa.ssa", func(){
	ContentType("application/json")

	Attributes(func() {
		Attribute("user_name", String, "User Name")
		Attribute("password", String, "User Password")
		Attribute("mail", String, "User mail-address")
		Attribute("group_id", String, "Gourp ID")
		Attribute("Data", String, "Data name")
		Attribute("data_name", String, "Data name")
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
})
