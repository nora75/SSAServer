package ssa

import (
	ssaserver "SSAServer/gen/ssa_server"
	"context"
	"log"
	"regexp"
	"fmt"
	"math/rand"
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
    rs6Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
    rs6LetterIdxBits = 6
    rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
    rs6LetterIdxMax = 63 / rs6LetterIdxBits
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
// random generate group id
// if not exists group id return false(error)
// dirの作成
func (s *sSAServersrvc) Register(ctx context.Context, p *ssaserver.RegisterPayload) (res *ssaserver.SsaResult, err error) {
	res = &ssaserver.SsaResult{}
	s.logger.Print("sSAServer.Register")
	res.UserName = &p.UserName
	res.Mail = &p.Mail
	res.UserID = &Counter
	Counter++
	GroupName := "group-" + RandString(12)
	if p.GroupID == nil || strings.EqualFold(*p.GroupID, ""){
		res.GroupID = &GroupName
	} else {
		res.GroupID = p.GroupID
	}
	return res, nil
}

// SSAへのログイン
// TODO
// dbとの統合
func (s *sSAServersrvc) Login(ctx context.Context, p *ssaserver.LoginPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Login")
    r := regexp.MustCompile(`pass`)
	if r.MatchString(p.Password) {
		return false, fmt.Errorf("パスワードが不正です。")
	}
	return true, nil
}

// グループIDを変更する
// TODO
// dbとの統合
func (s *sSAServersrvc) ChangeGroup(ctx context.Context, p *ssaserver.ChangeGroupPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Change_group")
	r := regexp.MustCompile(`^group-`)
	if r.MatchString(p.Password) {
		return false , fmt.Errorf("不正なグループ名です。")
	}
	// oldpath := GetUserDirPath(p.GroupID, p.UserID)
	oldpath := GetUserDirPath("group-"+p.Password, p.UserID)
	newpath := GetUserDirPath(p.GroupID, p.UserID)
	err = MoveUserDir(oldpath, newpath)
	if err !=  nil {
		return false, err
	}
	return true , nil
}

// 既存ユーザーの消去
// TODO
// dbとの統合
func (s *sSAServersrvc) DeleteUser(ctx context.Context, p *ssaserver.DeleteUserPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Delete_user")
	r := regexp.MustCompile(`^group-`)
	if r.MatchString(p.Password) {
		return false, fmt.Errorf("パスワードが不正です。")
	}
	path := GetUserDirPath("group-"+p.Password, p.UserID)
	err = DeleteUserDir(path)
	if err !=  nil {
		return false, err
	}
	// err = DeleteGroupDir(path)
	// if err !=  nil {
	// 	fmt.Println("An error occured when delete group dir:"+path)
	// }
	return true, nil
}

// データをサーバーへ保存する
// TODO
// dbとの統合
func (s *sSAServersrvc) SaveData(ctx context.Context, p *ssaserver.SaveDataPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Save_data")
	var mes string
	switch p.DataType {
		case 0: mes = "日記"
		case 1: mes = "録音"
		default: return false, fmt.Errorf("不正なdata_typeです。")
	}
	path := GetUserDirPath(p.GroupID, p.UserID)
	err = SaveFile(p.Data, path, p.DataName)
	if err !=  nil {
		return false, err
	}
	s.logger.Print(mes + ":" + p.DataName + "が保存されました。")
	if !(p.Image == nil && p.ImageName == nil || strings.EqualFold(*p.ImageName, "")){
		err = SaveFile(p.Image, path, *p.ImageName)
		if err !=  nil {
			return false, err
		}
		s.logger.Print(mes + ":" + *p.ImageName + "が保存されました。")
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
	fmt.Println("GroupID:"+p.GroupID)
	fmt.Println("DataUserID:",p.DataUserID)
	fmt.Println("DataType:",p.DataType)
	fmt.Println("UserID:",p.UserID)
	fmt.Println("DataName:",p.DataName)
	res.DataName = &p.DataName
	if !(p.ImageName == nil || strings.EqualFold(*p.ImageName, "")){
		fmt.Println("ImageName:",p.ImageName)
		res.ImageName = p.ImageName
	}
	return res, nil
}

// RandString Return Random n length String
func RandString(n int) string {
    b := make([]byte, n)
    cache, remain := randSrc.Int63(), rs6LetterIdxMax
    for i := n-1; i >= 0; {
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
