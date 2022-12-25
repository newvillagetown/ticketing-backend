package productNotice

import "main/common/noticeCommon"

type ProductRegisterNotice struct {
	ProductID  string `json:"product_id"`
	Name       string `json:"name"`
	PerAmount  string `json:"per_amount"`
	TotalCount string `json:"total_count"`
	RestCount  string `json:"rest_count"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

func (p *ProductRegisterNotice) Send() error {
	// 새로운 디이스 등록을 완료했습니다.
	title := "새로운 제품 등록이 됐습니다."
	text := "제품 ID : " + p.ProductID + "\n" +
		"상품 이름 : " + p.Name + "\n" +
		"시작 날짜 : " + p.StartDate + "\n" +
		"종료 날짜 : " + p.EndDate + "\n" +
		"티켓 금액 : " + p.PerAmount + "\n" +
		"티켓 총 수량 : " + p.TotalCount + "\n" +
		"남은 티켓 수량 : " + p.RestCount + "\n"
	widgets := make([]map[string]interface{}, 2)
	widgets[0] = map[string]interface{}{
		"textParagraph": map[string]string{
			"text": text,
		}}
	header := map[string]string{
		"title": title,
	}

	payload := map[string]interface{}{
		"cards": []map[string]interface{}{{
			"header": header,
			"sections": []map[string][]map[string]interface{}{{
				"widgets": widgets,
			}},
		}},
	}
	err := noticeCommon.GoogleChatSend(noticeCommon.GoogleChatUrl, payload)
	if err != nil {
		return err
	}

	return nil
}
