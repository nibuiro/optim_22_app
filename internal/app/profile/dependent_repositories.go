package profile

/*
 * profileサービスが依存しているリポジトリのハンドラ
 * 今回クリーンアーキテクチャもどきで組んでいる担当が私のみ
 * であるためここに記述を行う
 * 本当であればインポートと依存性注入でリポジトリハンドラが利用できることが望ましい
 */

import (
  "context"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
)


//#region request
type RequestRepository interface {
  Get(ctx context.Context, userId int) (profile, error)
}


type requestRepository struct {
  db     *gorm.DB
  logger log.Logger
}


func (r requestRepository) GetRequested(ctx context.Context, userId int) ([]typefile.Request, error) {
  var requesteds []typefile.Request
  result := r.db.WithContext(ctx).Find(&requesteds, "ClientID = ?", userId)
  return requesteds, result.Error
}
