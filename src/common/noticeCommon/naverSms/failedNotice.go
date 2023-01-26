package naverSms

import (
	"main/common/noticeCommon"
	"strconv"
)

type NaverSmsFailedNotice struct {
	PhoneList []string
	Content   string
}

func (n *NaverSmsFailedNotice) Send() error {
	//
	phoneList := ""
	for i := 0; i < len(n.PhoneList); i++ {
		phoneList = n.PhoneList[i] + "\n"
	}
	title := "NAVER SMS 메시지 전송 실패 리스트"
	text := "실패 총 인원 : " + strconv.Itoa(len(n.PhoneList)) + "\n" +
		"메시지 내용 : " + n.Content + "\n" +
		"휴대폰 번호 리스트 : " + phoneList + "\n"

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
