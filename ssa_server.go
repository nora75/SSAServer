package ssa

import (
	ssaserver "SSAServer/gen/ssa_server"
	"context"
	"log"
	"regexp"
	"fmt"
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
func (s *sSAServersrvc) Register(ctx context.Context, p *ssaserver.RegisterPayload) (res *ssaserver.SsaResult, err error) {
	res = &ssaserver.SsaResult{}
	s.logger.Print("sSAServer.Register")
	return res, nil
}

// SSAへのログイン
func (s *sSAServersrvc) Login(ctx context.Context, p *ssaserver.LoginPayload) (res *ssaserver.SsaResult, err error) {
	res = &ssaserver.SsaResult{}
	s.logger.Print("sSAServer.Login")
	return res, nil
}

// グループIDを変更する
func (s *sSAServersrvc) ChangeGroup(ctx context.Context, p *ssaserver.ChangeGroupPayload) (err error) {
	s.logger.Print("sSAServer.Change_group")
	return nil
}

// 既存ユーザーの消去
func (s *sSAServersrvc) DeleteUser(ctx context.Context, p *ssaserver.DeleteUserPayload) (err error) {
	s.logger.Print("sSAServer.Delete_user")
    r := regexp.MustCompile(`pass`)
	var mes string
	if r.MatchString(p.Password) {
		mes = "パスワードが不正です。"
	} else {
		mes = "正常に消去出来ました"
	}
	return fmt.Errorf("err %s", mes)
}

// データをサーバーへ保存する
func (s *sSAServersrvc) SaveData(ctx context.Context, p *ssaserver.SaveDataPayload) (err error) {
	s.logger.Print("sSAServer.Save_data")
	var mes string
	if 1 == *p.DataType {
		mes = "日記"
	} else {
		mes = "録音"
	}
	return fmt.Errorf("err %s", mes)
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
