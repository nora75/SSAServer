package ssa

import (
	ssaserver "SSAServer/gen/ssa_server"
	"context"
	"log"
	"regexp"
	"fmt"
)

// Counter is counter of registration
var Counter = 0
// GroupName is const of sample group name using return
var GroupName = "test-group"

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
func (s *sSAServersrvc) Register(ctx context.Context, p *ssaserver.RegisterPayload) (res *ssaserver.SsaResult, err error) {
	res = &ssaserver.SsaResult{}
	s.logger.Print("sSAServer.Register")
	res.UserName = &p.UserName
	res.Mail = &p.Mail
	res.UserID = &Counter
	Counter++
	if p.GroupID == nil {
		res.GroupID = &GroupName
	} else {
		res.GroupID = p.GroupID
	}
	return res, nil
}

// SSAへのログイン
func (s *sSAServersrvc) Login(ctx context.Context, p *ssaserver.LoginPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Login")
    r := regexp.MustCompile(`pass`)
	if r.MatchString(p.Password) {
		return false, fmt.Errorf("パスワードが不正です。")
	}
	return true, nil
}

// グループIDを変更する
func (s *sSAServersrvc) ChangeGroup(ctx context.Context, p *ssaserver.ChangeGroupPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Change_group")
    r := regexp.MustCompile(`group`)
	if r.MatchString(p.GroupID) {
		return true , nil
	}
	return false , fmt.Errorf("不正なグループ名です。")
}

// 既存ユーザーの消去
func (s *sSAServersrvc) DeleteUser(ctx context.Context, p *ssaserver.DeleteUserPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Delete_user")
    r := regexp.MustCompile(`pass`)
	if r.MatchString(p.Password) {
		return false, fmt.Errorf("パスワードが不正です。")
	}
	return true, nil
}

// データをサーバーへ保存する
func (s *sSAServersrvc) SaveData(ctx context.Context, p *ssaserver.SaveDataPayload) (res bool, err error) {
	s.logger.Print("sSAServer.Save_data")
	var mes string
	if 1 == *p.DataType {
		mes = "日記"
	} else if 2 == *p.DataType {
		mes = "録音"
	} else {
		return false, fmt.Errorf("不正なdata_typeです。")
	}
	return true, fmt.Errorf("%s", mes)
}

// データのリストを取得する
func (s *sSAServersrvc) ReturnDataList(ctx context.Context, p *ssaserver.ReturnDataListPayload) (res ssaserver.SsaResultCollection, view string, err error) {
	view = "data_list_origin"
	s.logger.Print("sSAServer.Return_data_list")
	return res, view, nil
}

// データをサーバーから取得する
func (s *sSAServersrvc) PickUpData(ctx context.Context, p *ssaserver.PickUpDataPayload) (res *ssaserver.SsaResult, err error) {
	res = &ssaserver.SsaResult{}
	s.logger.Print("sSAServer.Pick_up_data")
	return res, nil
}
