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
}


type repository struct {
  db     *gorm.DB
  logger log.Logger
}


func (r repository) Get(ctx context.Context, userId int) (typefile.Profile, error) {
  var userProfile typefile.Profile
  result := r.db.WithContext(ctx).Find("ID = ?", userId, &userProfile)
  if result.Error != nil {
    return typefile.Profile{}, result.Error
  } else {
    return userProfile, nil
  }
}

func (s repository) Create(ctx context.Context, userProfile *typefile.Profile) error {
  result := r.db.WithContext(ctx).Create(userProfile)
  return result.Error
}


func (s repository) Update(ctx context.Context, userProfile *typefile.Profile) error {
  /*
   * idがユニークであることによりupdateと等しい操作になることを期待
   * [MySQL ：： MySQL 5.6 リファレンスマニュアル ：： 13.2.5.3 INSERT ... ON DUPLICATE KEY UPDATE 構文]
   * (https://dev.mysql.com/doc/refman/5.6/ja/insert-on-duplicate.html)
   */
  result := r.db.WithContext(ctx).Create(userProfile)
  return result.Error
}


func (s repository) Delete(ctx context.Context, userId int) error {
  result := r.db.WithContext(ctx).Delete(&typefile.Profile{}, userId)
  return result.Error
}

//#region 実績を取得
func (r requestRepository) GetRequested(ctx context.Context, userId int) ([]typefile.Request, error) {
  var requesteds []typefile.Request
  result := r.db.WithContext(ctx).Find(&requesteds, "ClientID = ?", userId)
  return requesteds, result.Error
}


func (r requestRepository) GetParticipated(ctx context.Context, userId int) ([]typefile.Request, error) {
  //エンジニアIDがuserIdに一致するsubmissionのリクエストIDとリクエストのIDでサブミッションをリクエストに表結合して所望の値のリストを得る
  var participateds []typefile.Request //Submission  Request
  result := r.db.WithContext(ctx).
    Model(&typefile.Submission{}).
    Select("request.Finish, submission.UpdatedAt, request.ClientID, request.RequestName, request.Content, submissionID").
    Joins("INNER JOIN \"request\" ON submission.RequestID = request.ID").
    Where("submission.EngineerID = ?", userId).
    Scan(&participateds)

  return participateds, result.Error
}
//#endregion

func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}