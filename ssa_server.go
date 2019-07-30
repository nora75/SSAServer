package ssa

import (
	ssaserver "SSAServer/gen/ssa_server"
	"context"
	"log"
	"regexp"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"strconv"
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
// dir 移動 recursive
func (s *sSAServersrvc) ChangeGroup(ctx context.Context, p *ssaserver.ChangeGroupPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Change_group")
	r := regexp.MustCompile(`group`)
	if r.MatchString(p.GroupID) {
		return true , nil
	}
	return false , fmt.Errorf("不正なグループ名です。")
}

// 既存ユーザーの消去
// TODO
// dbとの統合
func (s *sSAServersrvc) DeleteUser(ctx context.Context, p *ssaserver.DeleteUserPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Delete_user")
	r := regexp.MustCompile(`pass`)
	if r.MatchString(p.Password) {
		return false, fmt.Errorf("パスワードが不正です。")
	}
	return true, nil
}

// データをサーバーへ保存する
// TODO
// dbとの統合
// mulipart
// fileの移動
// データをファイルと分ける
// ファイルを複数分ける
func (s *sSAServersrvc) SaveData(ctx context.Context, p *ssaserver.SaveDataPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Save_data")
	var mes string
	switch p.DataType {
	case 1: mes = "日記"
	case 2: mes = "録音"
	default: return false, fmt.Errorf("不正なdata_typeです。")
	}
	path := GetSavePath(p.GroupID, p.UserID)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0700)
		if err !=  nil {
			return false, err
		}
	}
	err = SaveFile(p.Data, path, p.DataName)
	if err !=  nil {
		return false, err
	}
	s.logger.Print(mes + ":" + p.DataName + "が保存されました。")
	if !(p.Image == nil && p.ImageName == nil || strings.EqualFold(*p.ImageName, "")){
		err = SaveFile(p.Image, path, *p.ImageName)
	}
	if err !=  nil {
		return false, err
	}
	s.logger.Print(mes + ":" + *p.ImageName + "が保存されました。")
	return true, nil
}

// データのリストを取得する
// TODO
// dbとの統合
// リストどうやって送信するの?
func (s *sSAServersrvc) ReturnDataList(ctx context.Context, p *ssaserver.ReturnDataListPayload) (res ssaserver.SsaResultCollection, view string, err error) {
	view = "data_list_origin"
	s.logger.Print("sSAServer.Return_data_list")
	return res, view, nil
}

// データをサーバーから取得する
// TODO
// dbとの統合
// データの実際の送信
// データのread
func (s *sSAServersrvc) PickUpData(ctx context.Context, p *ssaserver.PickUpDataPayload) (res *ssaserver.SsaResult, err error) {
	res = &ssaserver.SsaResult{}
	s.logger.Print("sSAServer.Pick_up_data")
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

// SaveFile save file in server
func SaveFile(data []byte,path string, name string) error {
	name = path + Getslash() + name
	f, err := os.Create(name)
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	return nil
}


// GetSavePath return current dir path
func GetSavePath(gr string, id int) (path string){
	path, _ = os.Getwd()
	slash := Getslash()
	path = path + slash + gr + slash + strconv.Itoa(id)
	return path
}

// Getslash return backslash or slash on server's environment
func Getslash() (slash string) {
	if runtime.GOOS == "windows" {
		slash = "\\\\"
	} else {
		slash = "/"
	}
	return slash
}
