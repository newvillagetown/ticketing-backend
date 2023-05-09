package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/common/awsCommon/ssm"
	"main/common/envCommon"
	_interface "main/features/system/test/usecase/interface"
	"net/http"
	"time"
)

type AuthTestUseCase struct {
	Repository     _interface.IAuthTestRepository
	ContextTimeout time.Duration
}

func NewAuthTestUseCase(repo _interface.IAuthTestRepository, timeout time.Duration) _interface.IAuthTestUseCase {
	return &AuthTestUseCase{Repository: repo, ContextTimeout: timeout}
}

// TODO 로직 정리가 필요
func (s *AuthTestUseCase) AuthTest(c context.Context) error {
	_, cancel := context.WithTimeout(c, s.ContextTimeout)
	defer cancel()

	// API 엔드포인트 URL 설정
	url := "https://msggw-auth.supersms.co:9440/auth/v1/token"

	// HTTP 요청 생성
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		// 오류 처리
		return err
	}
	id, _ := ssm.AwsGetParam(fmt.Sprintf("infobank_client_id_%s", envCommon.Env.Env))
	pw, _ := ssm.AwsGetParam(fmt.Sprintf("infobank_client_pw_%s", envCommon.Env.Env))
	// 헤더 설정
	req.Header.Set("X-IB-Client-Id", id)
	req.Header.Set("X-IB-Client-Passwd", pw)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json")
	// HTTP 클라이언트 생성
	var httpClient = http.DefaultClient
	// HTTP 요청 실행
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// HTTP 응답 처리
	responseBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var responseBody map[string]interface{}
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(responseBody["accessToken"].(string))
	return nil
}
