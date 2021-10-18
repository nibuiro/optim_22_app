# テスト計画

* 重要度

   service > api > repository

* 理由
  * serviceは値の検証、型変換、複数のリポジトリへの問い合わせと統合、apiに渡すなど複雑度が高くなりがち。
  * apiは条件分岐によるステータスコードに気を払う必要がある。
  * 結論としてrepositoryのテストは実際のDBサーバと正常なIOがなせることを確認すべき。go-sqlmockを用いてテストを書いていたが最も意味のないテストであった。むしろgo-sqlmockの作法に合わせてテストコードを書くなど目的と手段が逆転した。

## api

時間的、複雑度的に見送る

## service

検証プログラムのテストを行う 

TestRegistrationInformationValidate/invalid_email_type:_invalid_hyphen

は失敗するが見送る。またドットから始まるメールアドレスについても考慮すべきだが、{'.','-'}の許可及びアットマークとTLDの強制のみと100文字制限のみと現在はなっている。

## repository

当初go-sqlmockを用いたテストを書いていた。削除。