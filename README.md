# SSAServer

The server for SSA Project that written by golang  
別リポジトリのSSAプロジェクトのサーバーサイドアプリケーションです。  

## 実行方法(How to use)
事前にGOPATH等は設定されている事が前提とないｒます。  

```sh
go get "gopkg.in/natefinch/lumberjack.v2"
export GO111MODULE=on
go mod init
go get -u "goa.design/goa/v3"
go get -u "goa.design/goa/v3/..."
go build ./cmd/ssa_server
./ssa_server(.exe)
```

## フレームワーク
- Goa
