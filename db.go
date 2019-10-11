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

func insertUser(userData User, db *gorm.DB) {
	for _, user := range userData {
		db.NewRecord(user)
		db.Create(&user)
	}
}

func insertData(dataData Data, db *gorm.DB) {
	for _, data = range dataData {
		db.NewRecord(data)
		db.Create(&data)
	}
}

func insertUserData(UserName, PassWord, Mail string, GroupID *string) {
	userData := User{user_name: UserName, password: PassWord, mail: Mail, group_id: *GroupID}

	insertUser(userData, db)
}

func insertDataData(UserID, DataType int, GroupID, DataName, ImageName string, Title *string) {
	dataData := Data{group_id: GroupID, id: UserID, title: *Title, data_name: DataName, image_name: ImageName, data_type: DataType}

	insertData(dataData, db)
}
