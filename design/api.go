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
		Attribute("user_id", MapOf(String, Int), "User id", func(){
			Example("The first user id", func(){
				Description("The user who has register first")
				Value(map[string]int{"id": 1})
			})
			Example("The 234's user id", func(){
				Description("The user who has register 234's")
				Value(map[string]int{"id": 234})
			})
		})
		Attribute("user_name", MapOf(String, String), "User Name", func(){
			Example("The user's name Hirano", func(){
				Description("Simple usecase")
				Value(map[string]string{"user_name": "Hirano"})
			})
			Example("The user's name Hirano Tatsuya", func(){
				Description("Can contain space")
				Value(map[string]string{"user_name": "Hirano Tatsuya"})
			})
			Example("The user's name 平野 竜也", func(){
				Description("Can use utf-8")
				Value(map[string]string{"user_name": "平野 竜也"})
			})
		})
		Attribute("password", MapOf(String, String), "User Password", func(){
			Example("The password is 'password'", func(){
				Description("The password is 'password' but shoudn't use it")
				Value(map[string]string{"password": "password"})
			})
			Example("The password is 'PassW0rd!!@'", func(){
				Description("The password is 'PassW0rd!!@'")
				Value(map[string]string{"password": "PassW0rd!!@"})
			})
		})
		Attribute("mail", MapOf(String, String), "User mail-address", func(){
			Example("Example of mail-address", func(){
				Description("Example of mail-address")
				Value(map[string]string{"mail": "hoge@hoge.com"})
			})
		})
		Attribute("group_id", MapOf(String, String), "Gourp ID", func(){
			Example("The group name is 'bzuiphjjgdas'", func(){
				Description("The user who belong to Group 'bzuiphjjgdas'")
				Value(map[string]string{"gourp_id": "bzuiphjjgdas"})
			})
		})
		Attribute("data_name", MapOf(String, String), "Data name", func(){
			Example("The example of record data", func(){
				Description("The record data that's name is 'Record_12_2019-06-02_12-07-35'")
				Value(map[string]string{"data_name": "Record_12_2019-06-02_12-07-35"})
			})
			Example("The example of Diary data", func(){
				Description("The diary data that's name is 'diary_312_2019-03-02_12-07-35'")
				Value(map[string]string{"data_name": "Record_312_2019-03-02_12-07-35"})
			})
		})
		Attribute("Data", Any, "Data")
		Attribute("data_type", MapOf(String, Int), "Data's name", func(){
			Example("The record", func(){
				Description("The data_type 0(record)")
				Value(map[string]int{"data_type": 0})
			})
			Example("The Diary", func(){
				Description("The data_type 1(diary)")
				Value(map[string]int{"data_type": 1})
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

	View("login_extended", func() {
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

