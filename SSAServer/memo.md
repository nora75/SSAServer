エンドポイントで実行される。ssa_server.goに自分で記述した関数が呼ばれる。
その結果がresに入る。resは自分で定義したものが元のレスポンスの型
そのresがエラーチェックされる。
エラー皆無なら、resがencodeResponseにてエンコード関数にctx(context), w(ResponseWriter), res(上記の結果入った奴)が代入され呼び出される。
encodeResponseは各種別でencode_decodeにて定義されている。
resの部分はv interfaceにて実装されている。
res := v.(*ssaserverviews.SsaResult)にてresが取り出され、それを元に、bodyが作成される。
wにヘッダが設定される
最終的に、エンコーダーにbodyが代入され呼び出され、その戻り値が返される。


各エンドポイントのハンドラによって呼び出される順番  

- r : サーバーへのHTTPリクエスト
- ctx : コンテキスト(何が入ってるかは把握出来ていない)
- w : HTTPのリスポンスライター(これに値を入れて、これでクライアントに送信)

1. payload = リクエストのdecode(r)
1. エラーチェック
1. res = リクエストを取り扱うssa_server.goの関数(payload)
1. エラーチェック
1. リクエストのdecode(ctx, w, res)
1. エラーチェック
