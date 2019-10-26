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
	ID         uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// User struct define users table's struct
type User struct {
	gorm.Model
	UserName string
	Password  string
	Mail      string
	GroupID  string
}

// Data struct define data table's struct
type Data struct {
	gorm.Model
	UserID int `grom:"primary_key:true;association_foreignkey:ID"`
	// UserID string `grom:"primary_key:true;association_jointable_foreignkey:ID"`
	GroupID   string
	DataName  string `gorm:"primary_key"`
	ImageName string
	Title      string
	DataType  int
}

// DB Connection Information
const (
	// Dialect define DB type
	Dialect = "mysql"

	// DBUser define connected user name
	DBUser = "root"

	// DBPass define DB User's password
	DBPass = ""

	// DBProtocol define connect method
	DBProtocol = "tcp(127.0.0.1:3306)"

	// DBName define connect DB name
	DBName = "SSADB"
)

func main() {
	db := connectGorm()
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{},&Data{})
}

// connect to db
func connectGorm() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect+"?parseTime=true")

	if err != nil {
		log.Println(err)
	}

	db.LogMode(true)

	return db
}

// insert data to users tale
func insertUser(userData User) error {
	db := connectGorm()
	defer db.Close()

	row := User{} // 構造体インスタンス化
	row.UserName = userData.UserName
	row.Password = userData.Password
	row.Mail = userData.Mail
	row.GroupID = userData.GroupID

	flag := db.NewRecord(row)
	if !flag {
		return fmt.Errorf("Can't create new record")
	}

	result := db.Create(&row)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// insert data to data table
func insertData(dataData Data) error {
	db := connectGorm()
	defer db.Close()

	row := Data{}
	row.UserID = dataData.UserID
	row.GroupID = dataData.GroupID
	row.DataName = dataData.DataName
	row.ImageName = dataData.ImageName
	row.Title = dataData.Title
	row.DataType = dataData.DataType

	result := db.Create(&row)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// InsertUserData insert user data to users table
func InsertUserData(UserName string, PassWord string, Mail string, GroupID string) error {
	userData := User{
		UserName: UserName,
		Password:  PassWord,
		Mail:      Mail,
		GroupID:  GroupID,
	}

	err := insertUser(userData)
	if err != nil {
		return err
	}
	return nil
}

// InsertDataData insert data data to data table
func InsertDataData(UserID int, GroupID string, DataName string, ImageName string, Title string, DataType int) error {
	dataData := Data{
		UserID:   UserID,
		GroupID:   GroupID,
		Title:      Title,
		DataName:  DataName,
		ImageName: ImageName,
		DataType:  DataType,
	}

	err := insertData(dataData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser delete user on database
func DeleteUser(UserID int, PassWord string) error {
	db := connectGorm()
	defer db.Close()

	err := PasswordAuthentication(UserID, PassWord)
	if err != nil {
		return err
	}
	var user User
	result := db.Where("id = ?", UserID).Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateGroupID update database's user's GroupID
func UpdateGroupID(UserID int, GroupID, PassWord string) error {
	db := connectGorm()
	defer db.Close()

	err := PasswordAuthentication(UserID, PassWord)
	if err != nil{
		return err
	}

	var user User
	result := db.Model(&user).Where("id = ?", UserID).Update("group_id", GroupID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindAllDataInGroup return all data information of GroupID
func FindAllDataInGroup(GroupID string) ([]string, error) {
	db := connectGorm()
	defer db.Close()

	var data []Data
	result := db.Where("group_id = ?", GroupID).Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	retData := retDataList(data)

	return retData, nil
}

// FindData return a data information of DataID
func FindData(DataID int) ([]string, error) {
	db := connectGorm()
	defer db.Close()

	var data []Data
	result := db.Where("id = ?", DataID).First(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	retData := retDataList(data)

	return retData, nil
}

// PasswordAuthentication return success authentication of password
func PasswordAuthentication(UserID int, PassWord string) error {
	db := connectGorm()
	defer db.Close()

	var user User
	db.Where("id = ?", UserID).First(&user)

	if PassWord != user.Password {
		fmt.Printf("PW Auth Failed...")

		return fmt.Errorf("PW Auth Failed")
	}
	fmt.Printf("PW Auth success !!")
	return nil
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
