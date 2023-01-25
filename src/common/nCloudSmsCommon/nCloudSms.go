package nCloudSmsCommon

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/common/awsCommon/ssm"
	"main/common/errorCommon"
	"net/http"
	"strconv"
	"time"
)

var nSmsAccessKey string
var nSmsSecretKey []byte
var nSmsSendUrlFull string
var nSmsSendUrl string
var nSmsSenderNumber string
var httpClient = http.DefaultClient

type SmsType uint8

func InitNSms() error {

	smsInfo, iErr := ssm.AwsGetParams([]string{
		"nCloud_access_key",
		"nCloud_secret_key",
		"sms_service_id",
		"sms_sender_number"})
	if iErr != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), "nCloud aws ssm get failed", errorCommon.ErrFromAwsSsm)
	}
	nSmsAccessKey = smsInfo[0]
	nSmsSecretKey = []byte(smsInfo[1])
	nSmsSendUrl = fmt.Sprintf("/sms/v2/services/%s/messages", smsInfo[2])
	nSmsSendUrlFull = fmt.Sprintf("https://sens.apigw.ntruss.com%s", nSmsSendUrl)
	nSmsSenderNumber = smsInfo[3]

	return nil
}
func NSmsSend(phoneNumber, smsType, contentType, content string) (map[string]interface{}, error) {
	fmt.Println(phoneNumber)
	if phoneNumber == "01029121993" {
		return nil, errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), fmt.Sprintf("fail to create SMS send body - %+v", "test"), errorCommon.ErrFromInternal)
	}
	ctx := context.TODO()
	body := map[string]interface{}{
		"type":        smsType,
		"contentType": contentType,
		"countryCode": "82",
		"from":        nSmsSenderNumber,
		"messages": []map[string]interface{}{
			{"to": phoneNumber},
		},
		"content": content,
	}
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return nil, errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), fmt.Sprintf("fail to create SMS send body - %+v", body), errorCommon.ErrFromInternal)
	}

	sigTimestamp := time.Now().UnixNano() / 1000000
	sigRaw := fmt.Sprintf("POST %s\n%d\n%s", nSmsSendUrl, sigTimestamp, nSmsAccessKey)
	mac := hmac.New(sha256.New, nSmsSecretKey)
	mac.Write([]byte(sigRaw))
	sigKey := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	ctxHttp, ctxHttpCancel := context.WithTimeout(ctx, time.Second*8)
	defer ctxHttpCancel()
	req, err := http.NewRequestWithContext(ctxHttp, http.MethodPost, nSmsSendUrlFull, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), "fail to create SMS send request", errorCommon.ErrFromInternal)
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("x-ncp-apigw-timestamp", strconv.FormatInt(sigTimestamp, 10))
	req.Header.Add("x-ncp-iam-access-key", nSmsAccessKey)
	req.Header.Add("x-ncp-apigw-signature-v2", sigKey)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, errorCommon.ErrorMsg(errorCommon.ErrPartner, errorCommon.Trace(), fmt.Sprintf("SMS send http call fail - %s", nSmsSendUrl), errorCommon.ErrFromNaver)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errorCommon.ErrorMsg(errorCommon.ErrPartner, errorCommon.Trace(), fmt.Sprintf("SMS send http response parse fail - %s", nSmsSendUrl), errorCommon.ErrFromNaver)
	}
	resBodyMap := make(map[string]interface{})
	if err := json.Unmarshal(resBody, &resBodyMap); err != nil {
		return nil, errorCommon.ErrorMsg(errorCommon.ErrPartner, errorCommon.Trace(), fmt.Sprintf("fail to parse SMS send response - %s", string(resBody)), errorCommon.ErrFromNaver)
	}
	if resCode, ok := resBodyMap["statusCode"]; !ok || resCode != "202" {
		return nil, errorCommon.ErrorMsg(errorCommon.ErrPartner, errorCommon.Trace(), fmt.Sprintf("seems fail to send SMS - %s", string(resBody)), errorCommon.ErrFromNaver)
	}
	return resBodyMap, nil
}
