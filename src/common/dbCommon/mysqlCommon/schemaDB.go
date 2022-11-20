package mysqlCommon

import "time"

type User struct {
	ID        string    `json:"id"`        //pk
	Name      string    `json:"name"`      //이름
	Email     string    `json:"email"`     //이메일
	Created   time.Time `json:"created"`   //생성 날짜
	IsDeleted bool      `json:"IsDeleted"` //활동 여부
}

type UserAuth struct {
	ID         string    `json:"id"`         //pk
	Provider   string    `json:"provider"`   //OAuth 제공자
	UserID     string    `json:"userID"`     //외래키
	LastSignIn time.Time `json:"lastSignIn"` //마지막 로그인
	Created    time.Time `json:"created"`    //생성 날짜
	IsDeleted  bool      `json:"isDeleted"`  //활동 여부
}
