package profile

import (
  "context"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
)


type Repository interface {
  Get(ctx context.Context, userId int) (typefile.Profile, error)
  Create(ctx context.Context, userProfile *typefile.Profile) error
  Update(ctx context.Context, userProfile *typefile.Profile) error
  Delete(ctx context.Context, userId int) error

  GetRequested(ctx context.Context, userId int) ([]typefile.Request, error)
  GetParticipated(ctx context.Context, userId int) ([]typefile.Request, error)
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}



func NewRepository(db *gorm.DB, logger log.Logger) Repository {
  return repository{db, logger}
}


func (r repository) Get(ctx context.Context, userId int) (typefile.Profile, error) {
  var userProfile typefile.Profile
  result := r.db.WithContext(ctx).Find(&userProfile, "ID = ?", userId)
  if result.Error != nil {
    return typefile.Profile{}, result.Error
  } else {
    return userProfile, nil
  }
}


func (r repository) Create(ctx context.Context, userProfile *typefile.Profile) error {
  result := r.db.WithContext(ctx).Create(userProfile)
  return result.Error
}


func (r repository) Update(ctx context.Context, userProfile *typefile.Profile) error {
  /*
   * idがユニークであることによりupdateと等しい操作になることを期待
   * [MySQL ：： MySQL 5.6 リファレンスマニュアル ：： 13.2.5.3 INSERT ... ON DUPLICATE KEY UPDATE 構文]
   * (https://dev.mysql.com/doc/refman/5.6/ja/insert-on-duplicate.html)
   */
  result := r.db.WithContext(ctx).Create(userProfile)
  return result.Error
}


func (r repository) Delete(ctx context.Context, userId int) error {
  result := r.db.WithContext(ctx).Delete(&typefile.Profile{}, userId)
  return result.Error
}

//#region 実績を取得
func (r repository) GetRequested(ctx context.Context, userId int) ([]typefile.Request, error) {
  var requesteds []typefile.Request
  result := r.db.WithContext(ctx).Find(&requesteds, "client_id = ?", userId)
  return requesteds, result.Error
}


func (r repository) GetParticipated(ctx context.Context, userId int) ([]typefile.Request, error) {
  //エンジニアIDがuserIdに一致するsubmissionのリクエストIDとリクエストのIDでサブミッションをリクエストに表結合して所望の値のリストを得る
  var participateds []typefile.Request //Submission  Request
  result := r.db.WithContext(ctx).
    Model(&typefile.Submission{}).
    Select("requests.finish, submissions.updated_at, requests.client_id, requests.request_name, requests.content, submissions.id").
    Joins("INNER JOIN requests ON submissions.request_id = requests.id").
    Where("submissions.engineer_id = ?", userId).
    Scan(&participateds)

  return participateds, result.Error
}
//#endregion

func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}