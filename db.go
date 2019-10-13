package ssa

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// コピペで構成されたクソコード :(

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

	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})
}

func connectGorm() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect)

	if err != nil {
		log.Println(err.Error())
	}

	return db
}

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
	user_name string `gorm:"size:255"`
	password  string `gorm:"size:255"`
	mail      string `gorm:"size:255"`
	group_id  string `gorm:"size:255"`
}

type Data struct {
	gorm.Model
	data_type  int
	group_id   string `gorm:"size:255"`
	data_name  string `gorm:"size:255"`
	image_name string `gorm:"size:255"`
	title      string `gorm:"size:255"`
}

func insertUser(userData User) {
	db := connectGorm()
	defer db.Close()

	row := User{} // 構造体インスタンス化
	row.user_name = userData.user_name
	row.password = userData.password
	row.mail = userData.mail
	row.group_id = userData.group_id

	db.Create(&row)
}

func insertData(dataData Data) {
	db := connectGorm()
	defer db.Close()

	row := Data{}
	row.data_type = dataData.data_type
	row.group_id = dataData.group_id
	row.data_name = dataData.data_name
	row.image_name = dataData.image_name
	row.title = dataData.title

	db.Create(&row)
}

func insertUserData(UserName, PassWord, Mail string, GroupID *string) {
	userData := User{
		user_name: UserName,
		password:  PassWord,
		mail:      Mail,
		group_id:  *GroupID,
	}

	insertUser(userData)
}

func insertDataData(GroupID, DataName, ImageName string, Title *string, DataType int) {
	dataData := Data{
		group_id:   GroupID,
		title:      *Title,
		data_name:  DataName,
		image_name: ImageName,
		data_type:  DataType,
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

// func findAll() []User{

// }

func passwordAuthentication(UserID int, PassWord string) bool {
	userStruct := retUserStruct(UserID)

	if PassWord == userStruct.password {
		return true
	}
	return false
}

func retUserStruct(UserID int) User {
	db := connectGorm()
	defer db.Close()

	var user User
	db.Where("id = ?", UserID).First(&user)

	return user
}

func retDataStruct(DataID int) Data {
	db := connectGorm()
	defer db.Close()

	var data Data
	db.Where("id = ?", DataID).First(&data)

	return data
}
