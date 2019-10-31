package ssa

import (
	ssaserver "SSAServer/gen/ssa_server"
	"context"
	"fmt"
	"log"
	"math/rand"
	db "SSAServer/db"
	files "SSAServer/files"

	// "os"
	"strings"
	"time"
)

// Counter is counter of registration
var Counter = 0

// contextKey define tokenContextKey's type
type contextKey string

// tokenContextKey define key type get group name in SaveData
const tokenContextKey contextKey = "net.TCPAddr"

// for RandString
var randSrc = rand.NewSource(time.Now().UnixNano())

// for RandString
const (
	rs6Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rs6LetterIdxBits = 6
	rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
	rs6LetterIdxMax  = 63 / rs6LetterIdxBits
)

// SSAServer service example implementation.
// The example methods log the requests and return zero values.
type sSAServersrvc struct {
	logger *log.Logger
}

// NewSSAServer returns the SSAServer service implementation.
func NewSSAServer(logger *log.Logger) ssaserver.Service {
	return &sSAServersrvc{logger}
}

// SSAへの新規登録
// TODO
func (s *sSAServersrvc) Register(ctx context.Context, p *ssaserver.RegisterPayload) (res *ssaserver.SsaResult, err error) {
	res = &ssaserver.SsaResult{}
	s.logger.Print("sSAServer.Register")
	res.UserName = &p.UserName
	res.Mail = &p.Mail
	GroupName := "group-" + RandString(12)
	if p.GroupID == nil || strings.EqualFold(*p.GroupID, "") {
		res.GroupID = &GroupName
	} else {
		res.GroupID = p.GroupID
	}
	var id int
	id, err = db.InsertUserData(*res.UserName, p.Password, *res.Mail, *res.GroupID)

	if err != nil {
		return res, err
	}
	res.UserID = &id

	err = files.CreateUserDir(*res.GroupID, *res.UserID)

	if err != nil {
		return res, err
	}

	return res, nil
}

// SSAへのログイン
// TODO
// dbとの統合
func (s *sSAServersrvc) Login(ctx context.Context, p *ssaserver.LoginPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Login")

	// ログイン情報の成否判定
	err = db.UserAuth(p.Mail, p.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}

// グループIDを変更する
// TODO
// dbとの統合
func (s *sSAServersrvc) ChangeGroup(ctx context.Context, p *ssaserver.ChangeGroupPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Change_group")
	var oldGroupID string
	oldGroupID, err = db.UpdateGroupID(p.UserID, p.GroupID, p.Password)
	if err != nil {
		return false, err
	}

	err = files.MoveUserDir(oldGroupID, p.GroupID, p.UserID)
	if err != nil {
		_, err = db.UpdateGroupID(p.UserID, oldGroupID, p.Password)
		if err != nil {
			return false, err
		}

		return false, err
	}
	return true, nil
}

// 既存ユーザーの消去
// TODO
// dbとの統合
func (s *sSAServersrvc) DeleteUser(ctx context.Context, p *ssaserver.DeleteUserPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Delete_user")
	// path := GetUserDirPath("group-"+p.Password, p.UserID)
	// err = DeleteUserDir(path)
	// if err != nil {
	// 	return false, err
	// }
	// err = DeleteGroupDir(path)
	// if err !=  nil {
	// 	fmt.Println("An error occured when delete group dir:"+path)
	// }

	err = db.DeleteUser(p.UserID, p.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}

// データをサーバーへ保存する
// TODO
// dbとの統合
func (s *sSAServersrvc) SaveData(ctx context.Context, p *ssaserver.SaveDataPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Save_data")
	var mes string
	switch p.DataType {
	case 0:
		mes = "録音"
	case 1:
		mes = "日記"
	default:
		return false, fmt.Errorf("不正なdata_typeです。")
	}
	err = db.InsertDataData(p.UserID, p.Password, p.GroupID, p.DataName, *p.ImageName, *p.Title, p.DataType)
	if err != nil {
		return false, err
	}

	err = files.SaveFile(p.Data, p.GroupID, p.UserID, p.DataName)
	if err != nil {
		return false, err
	}
	s.logger.Print(mes + ":" + p.DataName + "が保存されました。")
	if (p.DataType == 1) {
		if !(p.Image == nil && (p.ImageName == nil || strings.EqualFold(*p.ImageName, ""))) {
			err = files.SaveFile(p.Image, p.GroupID, p.UserID, *p.ImageName)
			if err != nil {
				return false, err
			}
			s.logger.Print(mes + ":" + *p.ImageName + "が保存されました。")
		}
	}

	return true, nil
}

// データのリストを取得する
// TODO
// dbとの統合
// リストどうやって送信するの?
// SsaResultを、SsaResultCollectionに入れるみたいだけども入れてみたけど出力(値が返却)されなかったよ?
func (s *sSAServersrvc) ReturnDataList(ctx context.Context, p *ssaserver.ReturnDataListPayload) (res ssaserver.SsaResultCollection, view string, err error) {
	s.logger.Print("sSAServer.Return_data_list")
	view = "data_list_origin"

	// res, err = findAllDataInGroup(p.GroupID)
	// 表示わからん

	return res, view, nil
}

// データをサーバーから取得する
// TODO
// dbとの統合
// データの実際の送信
// データのread
func (s *sSAServersrvc) PickUpData(ctx context.Context, p *ssaserver.PickUpDataPayload) (res *ssaserver.SsaResult, err error) {
	s.logger.Print("sSAServer.Pick_up_data")
	res = &ssaserver.SsaResult{}
	res.GroupID = &p.GroupID
	res.DataType = &p.DataType
	res.UserID = &p.UserID
	fmt.Println("GroupID:" + p.GroupID)
	fmt.Println("DataUserID:", p.DataUserID)
	fmt.Println("DataType:", p.DataType)
	fmt.Println("UserID:", p.UserID)
	fmt.Println("DataName:", p.DataName)
	res.DataName = &p.DataName
	if !(p.ImageName == nil || strings.EqualFold(*p.ImageName, "")) {
		fmt.Println("ImageName:", p.ImageName)
		res.ImageName = p.ImageName
	}

	// retData := findData(p.GroupID)
	// 表示わからん

	return res, nil
}

// RandString Return Random n length String
func RandString(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rs6LetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rs6LetterIdxMax
		}
		idx := int(cache & rs6LetterIdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i--
		}
		cache >>= rs6LetterIdxBits
		remain--
	}
	return string(b)
}