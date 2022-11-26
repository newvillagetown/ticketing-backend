package mysqlCommon

type User struct {
	ID        string `json:"id"`        //pk
	Name      string `json:"name"`      //이름
	Email     string `json:"email"`     //이메일
	Created   string `json:"created"`   //생성 날짜
	IsDeleted bool   `json:"IsDeleted"` //활동 여부
}

type UserAuth struct {
	ID         string `json:"id"`         //pk
	Provider   string `json:"provider"`   //OAuth 제공자
	UserID     string `json:"userID"`     //외래키
	LastSignIn string `json:"lastSignIn"` //마지막 로그인
	Created    string `json:"created"`    //생성 날짜
	IsDeleted  bool   `json:"isDeleted"`  //활동 여부
}

type Product struct {
	ID          string `json:"id"`          //pk
	Created     string `json:"created"`     //생성 날짜
	LastUpdated string `json:"lastUpdated"` //마지막 수정 날짜
	IsDeleted   bool   `json:"isDeleted"`   //삭제 유무
	Name        string `json:"name"`        // 상품이름
	Description string `json:"description"` //상품 설명
	Category    string `json:"category"`    //상품 카테고리
	PerAmount   int64  `json:"perAmount"`   //상품 티켓 당 금액
	ImgUrl      string `json:"imgUrl"`      //이미지
	TotalCount  int64  `json:"totalCount"`  //총 수량
	RestCount   int64  `json:"restCount"`   //남은 수량
}
