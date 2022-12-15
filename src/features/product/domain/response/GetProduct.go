package response

type ResGetProduct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`        // 상품이름
	Description string `json:"description"` //상품 설명
	Category    string `json:"category"`    //상품 카테고리
	PerAmount   int64  `json:"perAmount"`   //상품 티켓 당 금액
	Image       string `json:"image"`       //이미지
	TotalCount  int64  `json:"totalCount"`  //총 수량
	RestCount   int64  `json:"restCount"`   //남은 수량
	StartDate   int64  `json:"startDate"`   //예매 시작 날짜 epoch time
	EndDate     int64  `json:"endDate"`     //예매 종료 날짜 epoch time
}
