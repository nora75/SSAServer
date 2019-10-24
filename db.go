package ssa

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`,
// which could be embedded in your models
type Model struct {
	id         uint `gorm:"primary_key"`
	created_at time.Time
	updated_at time.Time
	deleted_at *time.Time `sql:"index"`
}

type User struct {
	gorm.Model
	User_name string
	Password  string
	Mail      string
	Group_id  string
}

type Data struct {
	gorm.Model
	Data_type  int
	Group_id   string
	Data_name  string
	Image_name string
	Title      string
}

const (
	// DB種別
	Dialect = "mysql"

	// ユーザー名
	DBUser = "root"

	// パスワード
	DBPass = ""

	// プロトコル
	DBProtocol = "tcp(127.0.0.1:3306)"

	// DB名
	DBName = "SSADB"
)

func main() {
	db := connectGorm()
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE = InnoDB")
	db.AutoMigrate(&User{})
}

func connectGorm() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect+"?parseTime=true")

	if err != nil {
		log.Println(err.Error())
	}

	db.LogMode(true)

	return db
}

func insertUser(userData User) {
	db := connectGorm()
	defer db.Close()

	row := User{} // 構造体インスタンス化
	row.User_name = userData.User_name
	row.Password = userData.Password
	row.Mail = userData.Mail
	row.Group_id = userData.Group_id

	db.NewRecord(row)
	db.Create(&row)
}

func insertData(dataData Data) {
	db := connectGorm()
	defer db.Close()

	row := Data{}
	row.Data_type = dataData.Data_type
	row.Group_id = dataData.Group_id
	row.Data_name = dataData.Data_name
	row.Image_name = dataData.Image_name
	row.Title = dataData.Title

	db.Create(&row)
}

func insertUserData(UserName, PassWord, Mail string, GroupID string) {
	userData := User{
		User_name: UserName,
		Password:  PassWord,
		Mail:      Mail,
		Group_id:  GroupID,
	}

	insertUser(userData)
}

func insertDataData(GroupID, DataName, ImageName string, Title string, DataType int) {
	dataData := Data{
		Group_id:   GroupID,
		Title:      Title,
		Data_name:  DataName,
		Image_name: ImageName,
		Data_type:  DataType,
	}

	insertData(dataData)
}

func deleteUser(UserID int, PassWord string) {
	db := connectGorm()
	defer db.Close()

	result := passwordAuthentication(UserID, PassWord)
	if result {
		var user User
		db.Where("id = ?", UserID).Delete(&user)
	}
}

func updateGroupID(UserID int, GroupID, PassWord string) {
	db := connectGorm()
	defer db.Close()

	result := passwordAuthentication(UserID, PassWord)
	if result {
		var user User
		db.Model(&user).Where("id = ?", UserID).Update("group_id", GroupID)
	}
}

func findAllDataInGroup(GroupID string) []string {
	db := connectGorm()
	defer db.Close()

	var data []Data
	db.Where("group_id = ?", GroupID).Find(&data)

	retData := retDataList(data)

	return retData
}

func findData(DataID int) []string {
	db := connectGorm()
	defer db.Close()

	var data []Data
	db.Where("id = ?", DataID).First(&data)

	retData := retDataList(data)

	return retData
}

func passwordAuthentication(UserID int, PassWord string) bool {
	db := connectGorm()
	defer db.Close()

	var user User
	db.Where("id = ?", UserID).First(&user)

	if PassWord == user.Password {
		fmt.Printf("PW Auth success !!")

		return true
	}
	fmt.Printf("PW Auth Failed...")
	return false
}

func retDataStruct(DataID int) Data {
	db := connectGorm()
	defer db.Close()

	var data Data
	db.Where("id = ?", DataID).First(&data)

	return data
}

func retDataList(data []Data) []string {
	var ret []string

	for _, row := range data {
		ret = append(ret, fmt.Sprintf("%+v\n", row))
	}

	return ret
}
