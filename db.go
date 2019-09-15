package ssa

import (
	"database/sql"
	"strconv"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// [警告]
// ORMを使わないくそこーど

// InsertUser は ユーザー を 追加する 関数 です
func InsertUser(UserName, PassWord, Mail string, GroupID *string) bool {
	db, err := sql.Open("mysql", "root:@/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	/* AUTO INCREMENTだからUser_id指定なし */
	_, err = db.Exec(`INSERT INTO User(user_name,password,mail,group_id,flag)VALUES("` + UserName + `","` + PassWord + `","` + Mail + `","` + *GroupID + `","` + strconv.Itoa(0) + `");`)
	if err != nil {
		panic(err.Error())
	}

	return true
}

// Authentication は ログイン を 認証する 関数 です
func Authentication(Mail, PassWord string) bool {
	db, err := sql.Open("mysql", "root:@/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec(`SELECT user_id FROM User WHERE mail = "` + Mail + `" AND password = "` + PassWord + `";`) //MailとPWが一致するID検索
	if err != nil {
		panic(err.Error())

		return false
	}

	return true
}

// UpdateGroupID は グループID を 変更する 関数 です
func UpdateGroupID(UserID int, Password, GroupID string) bool {
	db, err := sql.Open("mysql", "root:@/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	pw := RequestPassword(UserID)
	if Password == pw {
		rows, err := db.Query(`SELECT group_id FROM User WHERE group_id = "` + GroupID + `";`)

		var data string
		for rows.Next() {
			err := rows.Scan(&data)
			if err != nil {
				panic(err.Error())
			}
		}

		if data != "" {
			_, err = db.Exec(`UPDATE user SET group_id = "` + GroupID + `" WHERE user_id = ` + strconv.Itoa(UserID) + `;`)
			if err != nil {
				panic(err.Error())
			}
		}

		return true
	}

	return false
}

// DeleteUserData は ユーザー を 削除する 関数 です
func DeleteUserData(UserID int, Password string) bool {
	db, err := sql.Open("mysql", "root:@/testdb")
	defer db.Close()

	pw := RequestPassword(UserID)
	if Password == pw {
		_, err = db.Exec(`DELETE FROM User WHERE user_id = ` + strconv.Itoa(UserID) + `;`)
		if err != nil {
			panic(err.Error())
		}

		return true
	}

	return false
}

// InsertData は データ の 情報 を 保存する 関数 です
// data_type: 0 = 日記 , 1 = 音声
func InsertData(UserID, DataType int, GroupID, DataName, ImageName string, Title *string) bool {
	db, err := sql.Open("mysql", "root:@/testdb")
	defer db.Close()

	_, err = db.Exec(`INSERT INTO data(group_id,user_id,data_name,title,image_name,data_type)VALUES("` + GroupID + `",` + strconv.Itoa(UserID) + `,"` + DataName + `","` + *Title + `","` + ImageName + `",` + strconv.Itoa(DataType) + `);`) // いんさーと
	if err != nil {
		panic(err.Error())
	}

	return true
}

// AllDataList は データ の 一覧 を 取得する 関数 です
func AllDataList(GroupID string) (retList [5][6]string) { // Slice使うとエラー吐くからArray
	db, err := sql.Open("mysql", "root:@/testdb")
	defer db.Close()

	rows, err := db.Query(`SELECT user_id,date,title,data_name,image_name,data_type FROM Data WHERE group_id = "` + GroupID + `";`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var UserID, Date, Title, DataName, ImageName, DataType string // temp変数
	j := 0
	for rows.Next() {
		err := rows.Scan(&UserID, &Date, &Title, &DataName, &ImageName, &DataType) // 同じ名前の変数に入れる？
		if err != nil {
			panic(err.Error())
		}

		dataList := [6]string{UserID, Date, Title, DataName, ImageName, DataType} // リストの中のリスト

		for i := 0; i < 5; i++ {
			// fmt.Println(strconv.Itoa(j) + ":" + strconv.Itoa(i))
			retList[j][i] = dataList[i]
		}
		j++
	}
	return retList
}

// PickUpDBData は data_name から 日記データ を 取得する 関数 です
func PickUpDBData(DataName string) (retList []string) {
	db, err := sql.Open("mysql", "root:@/testdb")
	defer db.Close()

	rows, err := db.Query(`SELECT user_id,date,title,data_name,image_name FROM Data WHERE data_name = "` + DataName + `";`) // data_idで検索
	if err != nil {
		panic(err.Error())
	}

	var UserID, Date, Title, PickedDataName, ImageName string
	for rows.Next() {
		err := rows.Scan(&UserID, &Date, &Title, &PickedDataName, &ImageName)
		if err != nil {
			panic(err.Error())
		}
		retList = []string{UserID, Date, Title, PickedDataName, ImageName}
	}
	return retList
}

// RequestPassword は UserID から パスワード を 問い合わせる 関数 です
func RequestPassword(UserID int) (retPW string) {
	db, err := sql.Open("mysql", "root:@/testdb")
	defer db.Close()

	rows, err := db.Query(`SELECT password FROM user WHERE user_id = "` + strconv.Itoa(UserID) + `";`) // パスワード検索
	if err != nil {
		panic(err.Error())
	}

	var Password string
	for rows.Next() {
		err := rows.Scan(&Password)
		if err != nil {
			panic(err.Error())
		}
		retPW = Password
	}
	return retPW
}
