package profile

import (
  "regexp"
//  "regexp"
  "github.com/go-ozzo/ozzo-validation/v4"
//  "github.com/go-ozzo/ozzo-validation/v4/is"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
  "encoding/json"
  "strconv"
  "context"
  "optim_22_app/internal/app/profile/repository"
//  "optim_22_app/internal/app/profile/repository"
)

//#region プロフィール登録情報
type sns struct {
  Github     string          `json:"github"`
  Twitter    string          `json:"twitter"`
  Facebook   string          `json:"facebook"`
}

type profile struct {
  Id         int             `json:"user_id"`
  Email      string          `json:"email"`
  Name       string          `json:"username"`
  Bio        string          `json:"comment"`
  Sns        json.RawMessage `json:"sns"`
  Icon       string          `json:"icon"`
  Submission json.RawMessage `json:"submissions"`
  Request    json.RawMessage `json:"requests"`
}


func (m profile) Validate() error {
  return validation.ValidateStruct(&m,
    //is unsigned integer
    //validation.Field(&m.Id, validation.Match(regexp.MustCompile("\\d+"))),
    //is BIO
    validation.Field(&m.Bio, validation.Length(0, 4000)),
    //is BASE64 encoded image, limited to 2MB ([MB] 2 * 1.33 ~ 2.67) 
    validation.Field(&m.Icon, validation.Length(0, 2.67e+6)),
  )
}
//#endregion
//#region 資格情報
type RegistrationInformation struct {
  Id       int    `json:"user_id"`
  Name     string `json:"username"`
  Email    string `json:"email"`
  Password string `json:"password"`
}


func (m RegistrationInformation) Validate() error {
  return validation.ValidateStruct(&m,
    validation.Field(&m.Name, validation.Required, validation.Length(1, 128)),
    //is.Email@ozzo-validation/v4/isはテストケース`success#1`にてエラー
    //{'.','-'}の許可及びアットマークとTLDの強制、半角1文字以上100文字以下制限のみ。
    validation.Field(&m.Email, validation.Required, validation.Length(1, 100), validation.Match(regexp.MustCompile("[a-zA-Z]+[a-zA-Z0-9\\.\\-]+@[a-zA-Z0-9\\-]+\\.[a-zA-Z0-9\\-\\.]+"))),
    //is SHA256
    validation.Field(&m.Password, validation.Required, validation.Length(64, 64), validation.Match(regexp.MustCompile("[A-Fa-f0-9]{64}$"))),
  )
}
//#endregion

type Service interface {
  Get(ctx context.Context, req string) (profile, error)
  Put(ctx context.Context, reqProfile profile, reqUser RegistrationInformation) error
}


type service struct {
  repo   Repository
  logger log.Logger
}

//新たなプロフィール操作サービスを作成
func NewService(repo Repository, logger log.Logger) Service {
  return service{repo, logger}
}


func (s service) Get(ctx context.Context, req string) (profile, error) {
  //リクエスト文字列を数値型ユーザIDに変換
  //var userId int
  userId, err := strconv.Atoi(req)
  if err != nil {
    s.logger.Error(err)
    return profile{}, err
  }
  //該当ユーザのプロフィールを取得
  //var userProfileWithRecords profile

  if userProfile, err := s.repo.Get(ctx, userId); err != nil {
    return profile{}, err
  } else {
    if requesteds, err := s.repo.GetRequested(ctx, userId); err != nil {
      return profile{}, err
    } else {
      //s.logger.Debug(userIds)
      nRequesteds := len(requesteds)
      var userIds []int
      s.logger.Debug(nRequesteds)

      //参照されるuserIdを抽出
      for _, requested := range requesteds {
        for _, engineer := range requested.Engineers {
          userIds = append(userIds, engineer.ID)
        }
      }//
      //参照されたuserIdのプロフィールを取得
      if engineerProfiles, err := s.repo.GetProfiles(ctx, userIds); err != nil {
        s.logger.Error(err)
        return profile{}, err
      } else {
        //取得したプロフィールをuserIdをキーとするハッシュテーブルに落とし込む
        engineerProfilesTable := make(map[int]roundary.Profile)
        for _, engineerProfile := range engineerProfiles {
          engineerProfilesTable[engineerProfile.ID] = engineerProfile
        }
        //#region 依頼参加者のプロフィール項目を埋める
        for iRequested := 0; iRequested < nRequesteds; iRequested++ {
          requested := &requesteds[iRequested]
          nEngineers  := len(requested.Engineers)
          for iEngineer := 0; iEngineer < nEngineers; iEngineer++ {
            Engineer := &requested.Engineers[iEngineer]
            //requesteds[iRequested].Engineers[iEngineer]
            Engineer.Bio = engineerProfilesTable[Engineer.ID].Bio
            Engineer.Sns = string(engineerProfilesTable[Engineer.ID].Sns[:])
            Engineer.Icon = engineerProfilesTable[Engineer.ID].Icon
          }
        }
        //#endregion
        if requestedsText, err := json.Marshal(requesteds); err != nil {
          s.logger.Error(err)
          return profile{}, err
        } else {
          if submitteds, err := s.repo.GetSubmitted(ctx, userId); err != nil {
            s.logger.Error(err)
            return profile{}, err
          } else {
            if submittedsText, err := json.Marshal(submitteds); err != nil {
              s.logger.Error(err)
              return profile{}, err
            } else {
              //#region as userProfileWithRecords
              userProfile.Submission = submittedsText
              userProfile.Request = requestedsText
              //#endregion
              return userProfile, nil
            }
          }
        }
      }
    }
  }
}


func (s service) Put(ctx context.Context, reqProfile profile, reqUser RegistrationInformation) error {
  //SNS登録情報を読み込み
  sns := sns{}
  json.Unmarshal(reqProfile.Sns, &sns)
  //リクエストの値を検証
  if err := reqProfile.Validate(); err != nil {
    s.logger.Error(err)
    return err
  }
  //リクエストの値を検証
  if err := reqUser.Validate(); err != nil {
    s.logger.Error(err)
    return err
  }
  //クエリの値を定義
  profileUpdates := typefile.Profile{
    ID:   reqProfile.Id,
    Bio:  reqProfile.Bio,
    Sns:  reqProfile.Sns,
    Icon: reqProfile.Icon,
  }
  userUpdates := typefile.User{
    ID:       reqUser.Id,
    Name:     reqUser.Name,
    Email:    reqUser.Email,
    Password: reqUser.Password,
  }
  //資格情報とプロフィールを更新
  if err := s.repo.Update(ctx, &profileUpdates, &userUpdates); err != nil {
    s.logger.Error(err)
    return err
  } else {
    return nil
  }
}


