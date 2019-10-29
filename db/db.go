package db

import (
	"fmt"
	"time"
	"strconv"

	// ORM
	"github.com/jinzhu/gorm"
	// alias of mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`,
// which could be embedded in your models
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// User struct define users table's struct
type User struct {
	gorm.Model
	UserName string `gorm:"not null"`
	Password string `gorm:"not null"`
	Mail     string `gorm:"not null"`
	GroupID  string `gorm:"not null"`
}

// Data struct define data table's struct
type Data struct {
	gorm.Model
	UserID    int    `grom:"primary_key:true"`
	GroupID   string `gorm:"not null"`
	DataName  string `gorm:"primary_key"`
	ImageName string
	Title     string
	DataType  int `gorm:"not null"`
}

// Result struct define return data's list
type Result struct {
	UserID    int
	UserName  string
	GroupID   string
	DataName  string
	ImageName string
	Title     string
	DataType  int
}

// DB Connection Information
const (
	// Dialect define DB type
	Dialect = "mysql"

	// DBUser define connected user name
	DBUser = ""

	// DBPass define DB User's password
	DBPass = ""

	// DBProtocol define connect method
	DBProtocol = "tcp(127.0.0.1:3306)"

	// DBName define connect DB name
	DBName = ""
)

// check connection and auto create tables if not available
func main() {
	db, err := connectGorm()
	defer db.Close()

	if err != nil {
		return
	}

	// auto create tables if not available
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{}, &Data{})
}

// connect to db
func connectGorm() (*gorm.DB, error) {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect+"?parseTime=true")

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}

// insert data to users tale
func insertUser(userData User) (int, error) {
	db, err := connectGorm()
	defer db.Close()

	if err != nil {
		return 0, err
	}

	var users User
	result := db.Table("users").Where("mail = ?", userData.Mail).Scan(&users)
	if users.ID != 0  {
		return 0, fmt.Errorf("Already current mail address is registered")
	}

	row := User{} // 構造体インスタンス化
	row.UserName = userData.UserName
	row.Password = userData.Password
	row.Mail = userData.Mail
	row.GroupID = userData.GroupID

	flag := db.NewRecord(row)
	if !flag {
		return 0, fmt.Errorf("Can't create new record")
	}

	result = db.Create(&row)
	if result.Error != nil {
		return 0, result.Error
	}

	id := int(result.Value.(*User).Model.ID)

	return id, nil
}

// insert data to data table
func insertData(dataData Data) error {
	db, err := connectGorm()
	defer db.Close()

	if err != nil {
		return err
	}

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
func InsertUserData(UserName string, PassWord string, Mail string, GroupID string) (int, error) {
	userData := User{
		UserName: UserName,
		Password: PassWord,
		Mail:     Mail,
		GroupID:  GroupID,
	}

	id, err := insertUser(userData)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// InsertDataData insert data data to data table
func InsertDataData(UserID int, GroupID string, DataName string, ImageName string, Title string, DataType int) error {
	dataData := Data{
		UserID:    UserID,
		GroupID:   GroupID,
		Title:     Title,
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
	db, err := connectGorm()
	defer db.Close()
	if err != nil {
		return err
	}

	err = passwordAuthentication(UserID, PassWord)
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
func UpdateGroupID(UserID int, GroupID string, PassWord string) (oldgid string, err error) {
	db, err := connectGorm()
	defer db.Close()
	if err != nil {
		return "", err
	}

	oldgid, err = findGroupIDbyUserID(UserID)

	err = passwordAuthentication(UserID, PassWord)
	if err != nil {
		return
	}

	err = updateUserTableGroupID(db, UserID, GroupID)
	if err != nil {
		return
	}
	var datas Data
	result := db.Model(&datas).Where("user_id = ?", UserID).Update("group_id", GroupID)
	if result.Error != nil {
		if err != nil {
			return
		}
		_ = updateUserTableGroupID(db, UserID, oldgid)
		err = result.Error
		return
	}

	return
}

// FindAllDataInGroup return all data information of GroupID
func FindAllDataInGroup(GroupID string) ([]Result, error) {
	db, err := connectGorm()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var ret []Result
	subq := db.Select("*").Table("data").Where("group_id = ?", GroupID).SubQuery()
	res := db.Table("users").Select("t1.user_id, users.user_name, t1.group_id, t1.data_name, t1.image_name, t1.title, t1.data_type").Joins("inner join ? as t1 on t1.user_id = users.id", subq).Scan(&ret)
	if res.Error != nil {
		fmt.Println("Error")
		return nil, res.Error
	}

	return ret, nil
}

// findData return a data information of DataID
func findData(DataID int) ([]string, error) {
	db, err := connectGorm()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var data []Data
	result := db.Where("id = ?", DataID).First(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	retData := retDataList(data)

	return retData, nil
}

// UserAuth test user authentication by mail and password
func UserAuth(mail string, password string) error {
	id, err := findIDbyMail(mail)
	if err != nil {
		return err
	}
	err = passwordAuthentication(id, password)
	if err != nil {
		return err
	}
	return nil
}

// passwordAuthentication return success authentication of password
func passwordAuthentication(UserID int, PassWord string) error {
	db, err := connectGorm()
	defer db.Close()
	if err != nil {
		return err
	}

	var user User
	db.Where("id = ?", UserID).First(&user)

	if PassWord != user.Password {
		fmt.Printf("User: " + strconv.Itoa(UserID) + " PW Auth Failed")

		return fmt.Errorf("PW Auth Failed")
	}
	fmt.Printf("User: " + strconv.Itoa(UserID) + " PW Auth success !!")
	return nil
}

func retDataStruct(DataID int) Data {
	var data Data
	db, err := connectGorm()
	defer db.Close()
	if err != nil {
		return data
	}

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

// find id by mail on users table
func findIDbyMail(mail string) (int, error) {
	db, err := connectGorm()
	defer db.Close()
	if err != nil {
		return 0, err
	}

	var users User
	result := db.Table("users").Select("id").Where("mail = ?", mail).Scan(&users)
	if result.Error != nil {
		return 0, result.Error
	}
	id := int(users.Model.ID)
	return id, nil
}

// findGroupIDbyUserID Find group id by userID on users table
func findGroupIDbyUserID(UserID int) (string, error) {
	db, err := connectGorm()
	defer db.Close()
	if err != nil {
		return "", err
	}

	var users User
	result := db.Table("users").Select("group_id").Where("user_id = ?", UserID).Scan(&users)
	if result.Error != nil {
		return "", result.Error
	}
	gid := string(users.GroupID)
	return gid, nil
}

// updateUserTableGroupID update group id of user tale
func updateUserTableGroupID(db *gorm.DB, UserID int, GroupID string) error {
	var user User
	result := db.Model(&user).Where("id = ?", UserID).Update("group_id", GroupID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}