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
 *                V                                |
 *  relational model definition                    |
 *                                                 |
 *                |                                |
 *                |                                |
 *                |                                |
 *    +---------  |  --------repository------------------------+
 *                |                                A
 *                |                                |
 *                V                            roundary
 *
 *
 *
 */


type Repository interface {
  Get(ctx context.Context, userId int) (profile, error)
  Create(ctx context.Context, userProfile *typefile.Profile) error
  Update(ctx context.Context, userProfile *typefile.Profile, userCredntial *typefile.User) error
  Delete(ctx context.Context, userId int) error

  GetProfiles(ctx context.Context, userIds []int) ([]roundary.Profile, error)
  GetRequested(ctx context.Context, userId int) ([]roundary.Request, error)
  GetParticipated(ctx context.Context, userId int) ([]roundary.Request, error)
  GetSubmitted(ctx context.Context, userId int) ([]roundary.Submission, error)
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


func (r repository) Update(ctx context.Context, userProfile *typefile.Profile, userCredntial *typefile.User) error {

  tx := r.db.WithContext(ctx).Begin()
  //ユーザ登録以降基本的に編集を受け付ける形となり初期エントリが必要
  err := UpdateCredentialAndProfile(tx, userProfile, userCredntial)
  defer func() {
    if err != nil {
      tx.Rollback()
    } else {
      tx.Commit()
    }
  }()

  if err != nil {
    return err
  } else {
    return nil
  }
}


func (r repository) Delete(ctx context.Context, userId int) error {
  result := r.db.WithContext(ctx).Delete(&roundary.Profile{}, userId)
  return result.Error
}


func (r repository) GetProfiles(ctx context.Context, userIds []int) ([]roundary.Profile, error) {
  var userProfiles []roundary.Profile

  result := r.db.WithContext(ctx).
    Find(&userProfiles, "id IN ?", userIds)

  if result.Error != nil {
    return make([]roundary.Profile, 1), result.Error
  } else {
    return userProfiles, nil
  }
}

//#region 実績を取得
func (r repository) GetRequested(ctx context.Context, userId int) ([]roundary.Request, error) {
  var requesteds []roundary.Request
  result := r.db.WithContext(ctx).
    Preload("Engineers").
    Preload("Winner").
    Preload("Submission").
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


func (r repository) GetSubmitted(ctx context.Context, userId int) ([]roundary.Submission, error) {
  var submissions []roundary.Submission
  result := r.db.WithContext(ctx).
    Find(&submissions, "submissions.engineer_id = ?", userId)
  return submissions, result.Error
}

//#endregion

func UpdateCredentialAndProfile(tx *gorm.DB, userProfile *typefile.Profile, userCredntial *typefile.User) error {

  result := tx.Model(userProfile).
    Updates(map[string]interface{}{"bio": userProfile.Bio, "sns": userProfile.Sns, "icon": userProfile.Icon})
  
  if err := result.Error; err != nil {
    return err
  } else {
    //pass
  }

  result = tx.Model(userCredntial).
    Updates(map[string]interface{}{"email": userCredntial.Email, "password": userCredntial.Password})
  
  if err := result.Error; err != nil {
    return err
  } else {
    //pass
  }

  return nil

}




func StubNewRepository(args ...interface{}) Repository {return repository{nil, nil}}