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
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	gorm.Model
	UserName string `gorm:"size:255"`
	PassWord string `gorm:"size:255"`
	Mail     string `gorm:"size:255"`
	GroupID  string `gorm:"size:255"`
}

type Data struct {
	gorm.Model
	DataType  int
	GroupID   string `gorm:"size:255"`
	DataName  string `gorm:"size:255"`
	ImageName string `gorm:"size:255"`
	Title     string `gorm:"size:255"`
}

// UserID, DataType int, GroupID, DataName string, ImageName, Title *string

func (u User) String() string {
	return fmt.Sprintf("%s(%d)", u.UserName, u.PassWord, u.Mail, u.GroupID)
}

func (d Data) String() string {
	return fmt.Sprintf("%s(%d)", d.GroupID, d.Title, d.DataName, d.ImageName, d.DataType)
}
