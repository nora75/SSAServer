package design

import (
	. "goa.design/goa/v3/dsl"
)

// MyResultType is declation of use in result views
var MyResultType = ResultType("application/vnd.ssa.result", func(){
	ContentType("application/json")

	Attributes(func() {
		Attribute("user_id", Int, "User id", func(){
			Example("The first user id", func(){
				Description("The user who has register first")
				Value(1)
			})
			Example("The 234's user id", func(){
				Description("The user who has register 234's")
				Value(234)
			})
		})
		Attribute("user_name", String, "User Name", func(){
			Example("The user's name 平野 竜也", func(){
				Description("Can use utf-8")
				Value("平野 竜也")
			})
		})
		Attribute("password", String, "User Password", func(){
			Example("The password is 'password'", func(){
				Description("The password is 'password' but shoudn't use it")
				Value("password")
			})
			Example("The password is 'PassW0rd!!@'", func(){
				Description("The password is 'PassW0rd!!@'")
				Value("PassW0rd!!@")
			})
		})
		Attribute("mail", String, "User mail-address", func(){
			Example("Example of mail-address", func(){
				Description("Example of mail-address")
				Format("email")
				Value("hoge@hoge.com")
			})
		})
		Attribute("group_id", String, "Group ID", func(){
			Example("The group name is 'bzuiphjjgdas'", func(){
				Description("The user who belong to Group 'bzuiphjjgdas'")
				Value("bzuiphjjgdas")
			})
		})
		Attribute("data_name", String, "Data name", func(){
			Example("The example of record data", func(){
				Description("Record data of User 12,Time 2019-06-02_12-07-35")
				Value("Record_12_2019-06-02_12-07-35")
			})
			Example("The example of Diary data", func(){
				Description("Diary data of User 312,Time 2019-03-02_12-07-35")
				Value("Diary_312_2019-03-02_12-07-35")
			})
		})
		Attribute("Data", Any, "Data")
		Attribute("Image", Any, "Image data")
		Attribute("data_type", Int, "Data's name", func(){
			Example("The record", func(){
				Description("The data_type 0(record)")
				Value(0)
			})
			Example("The Diary", func(){
				Description("The data_type 1(diary)")
				Value(1)
			})
		})
		Attribute("title", String, "Data title", func(){
			Example("Title: hoge", func(){
				Description("The Diary's title of hoge")
				Value("hoge")
			})
		})
		Attribute("image_name", String, "image's name", func(){
			Example("Sample image name", func(){
				Description("The image's title of huge")
				Value("huge")
			})
		})
		Attribute("date_time", String, "date time", func(){
			Example("date time", func(){
				Description("uploaded time")
				Value("2019-06-05-02-02")
			})
		})

		View("default", func() {
			Attribute("user_id")
			Attribute("mail")
			Attribute("user_name")
			Example("example", func(){
				Description("example of default")
				Value(map[string]string{
					"user_id": "1",
					"user_name": "Hirano",
					"mail": "Hirano@test.com",
				})
			})
		})

		View("extended", func() {
			Attribute("user_id")
			Attribute("mail")
			Attribute("user_name")
			Attribute("group_id")
		})

		View("data", func() {
			Attribute("Data")
			Attribute("data_type")
			Attribute("data_name")
		})

		View("data_extended", func() {
			Attribute("Data")
			Attribute("data_type")
			Attribute("data_name")
			Attribute("title")
		})

		View("data_extended_with_image", func() {
			Attribute("Data")
			Attribute("data_type")
			Attribute("data_name")
			Attribute("title")
			Attribute("Image")
			Attribute("image_name")
		})

		View("data_list_origin", func() {
			Attribute("data_type")
			Attribute("data_name")
			Attribute("title")
			Attribute("date_time")
			Attribute("user_name")
		})

	})
})

// DataResult used by return data's list
var DataResult = CollectionOf(MyResultType, func(){
	View("data_list_origin")
})

// CustomErrorType is declation for responce error
var CustomErrorType = ResultType("application/vnd.ssa.error", func() {
	Attributes(func() {
		Attribute("fault", Boolean, "Return false when Error occur", func() {
			Default(true)
			Example(true)
		})
		Attribute("message", String, "Error returned.", func() {
			Meta("struct:error:name")
			Example("不正なリクエストです")
		})
		Required("fault", "message")
	})
	View("default", func() {
		Attribute("fault")
		Attribute("message")
	})
})

var _ = API("SSA", func() {
	Title("SSAServer")
	Description("SSAサービスのサーバーサイドプログラム")
	Version("0.1")
	Server("SSAServer", func() {
		Description("ローカルホストでのテスト段階")
		Services("SSAServer", "swagger")
		Host("localhost", func() {
			URI("http://localhost:8000/")
		})
	})
})

