package profile

import (
  "context"
  "gorm.io/gorm"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
  "optim_22_app/internal/app/profile/repository"
)

/*
 *              input                            output
 *
 *                |                                |
 *                |                                |
 *                u                                |
 *  relational model definition                    |
 *                                                 |
 *                |                                |
 *                |                                |
 *                |                                |
 *    +---------  |  --------repository------------------------+
 *                |                                n
 *                |                                |
 *                u                            roundary
 *
 *
 *
 */


type Repository interface {
  Get(ctx context.Context, userId int) (profile, error)
  Create(ctx context.Context, userProfile *typefile.Profile) error
  Update(ctx context.Context, userProfile *typefile.Profile) error
  Delete(ctx context.Context, userId int) error

  GetRequested(ctx context.Context, userId int) ([]roundary.Request, error)
  GetParticipated(ctx context.Context, userId int) ([]roundary.Request, error)
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}



func NewRepository(db *gorm.DB, logger log.Logger) Repository {
  return repository{db, logger}
}


func (r repository) Get(ctx context.Context, userId int) (profile, error) {
  var userProfile profile

  result := r.db.WithContext(ctx).
    Model(&roundary.Profile{}).
    Select("profiles.id, users.name, users.email, profiles.bio, profiles.sns, profiles.icon").
    Joins("INNER JOIN users ON profiles.id = users.id").
    Where("profiles.id = ?", userId).
    Scan(&userProfile)

  if result.Error != nil {
    return profile{}, result.Error
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
  result := r.db.WithContext(ctx).Delete(&roundary.Profile{}, userId)
  return result.Error
}

//#region 実績を取得
func (r repository) GetRequested(ctx context.Context, userId int) ([]roundary.Request, error) {
  var requesteds []roundary.Request
  result := r.db.WithContext(ctx).
    Preload("Engineers").
    Preload("Client").
    Preload("Winner").
    Find(&requesteds, "client_id = ?", userId)
  return requesteds, result.Error
}


func (r repository) GetParticipated(ctx context.Context, userId int) ([]roundary.Request, error) {
  //エンジニアIDがuserIdに一致するsubmissionのリクエストIDとリクエストのIDでサブミッションをリクエストに表結合して所望の値のリストを得る
  var participateds []roundary.Request //Submission  Request
  var client_ids []int
  result := r.db.Select("client_id").
    Joins("INNER JOIN requests ON submissions.request_id = requests.id").
    Where("submissions.engineer_id = ?", userId).
    Table("submissions").Scan(&client_ids)

  result = r.db.WithContext(ctx).
    Preload("Engineers").
    Preload("Client").
    Preload("Winner").
    Find(&participateds, "client_id IN ?", client_ids)

  return participateds, result.Error
}
//#endregion

func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}