// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v0.1/auth/google/signin": {
            "get": {
                "description": "■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "google 로그인",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            }
        },
        "/v0.1/auth/google/signin/callback": {
            "get": {
                "description": "■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "google login callback",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            }
        },
        "/v0.1/auth/google/signout": {
            "get": {
                "description": "■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "google 로그아웃",
                "parameters": [
                    {
                        "type": "string",
                        "description": "accessToken",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            }
        },
        "/v0.1/auth/user/withdrawal": {
            "post": {
                "description": "■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "회원 탈퇴",
                "parameters": [
                    {
                        "description": "json body",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqWithdrawalUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            }
        },
        "/v0.1/features/message/naver/sms": {
            "post": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\n\n■ errCode with 401\nTOKEN_BAD : 토큰 인증 실패\nPOLICY_VIOLATION : 토큰 세션 정책 위반\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "message"
                ],
                "summary": "네이버 sms 메시지 전송",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "phoneList",
                        "name": "phoneList",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "contentType",
                        "name": "contentType",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "smsType",
                        "name": "smsType",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "reserveTime",
                        "name": "reserveTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "scheduleCode",
                        "name": "scheduleCode",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "image",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            }
        },
        "/v0.1/features/product": {
            "get": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\n\n■ errCode with 401\nTOKEN_BAD : 토큰 인증 실패\nPOLICY_VIOLATION : 토큰 세션 정책 위반\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "상품 상세정보 가져오기",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productID",
                        "name": "productID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResGetProduct"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            },
            "put": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\n\n■ errCode with 401\nTOKEN_BAD : 토큰 인증 실패\nPOLICY_VIOLATION : 토큰 세션 정책 위반\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "상품 수정하기",
                "parameters": [
                    {
                        "description": "json body",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqUpdateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            },
            "post": {
                "description": "■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "상품 등록",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "category",
                        "name": "category",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "image",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "perAmount",
                        "name": "perAmount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "totalCount",
                        "name": "totalCount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "restCount",
                        "name": "restCount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "startDate",
                        "name": "startDate",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "endDate",
                        "name": "endDate",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            },
            "delete": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\n\n■ errCode with 401\nTOKEN_BAD : 토큰 인증 실패\nPOLICY_VIOLATION : 토큰 세션 정책 위반\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "상품 삭제하기(소프트)",
                "parameters": [
                    {
                        "description": "json body",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqDeleteProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            }
        },
        "/v0.1/features/product/gets": {
            "get": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\n\n■ errCode with 401\nTOKEN_BAD : 토큰 인증 실패\nPOLICY_VIOLATION : 토큰 세션 정책 위반\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "상품 목록 가져오기",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResGetsProduct"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorCommon.ResError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errorCommon.ResError": {
            "type": "object",
            "properties": {
                "errType": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "request.ReqDeleteProduct": {
            "type": "object",
            "required": [
                "productID"
            ],
            "properties": {
                "productID": {
                    "type": "string"
                }
            }
        },
        "request.ReqUpdateProduct": {
            "type": "object",
            "required": [
                "productID"
            ],
            "properties": {
                "category": {
                    "description": "상품 카테고리",
                    "type": "string"
                },
                "description": {
                    "description": "상품 설명",
                    "type": "string"
                },
                "endDate": {
                    "description": "예매 종료 날짜 epoch time",
                    "type": "integer"
                },
                "name": {
                    "description": "상품이름",
                    "type": "string"
                },
                "perAmount": {
                    "description": "상품 티켓 당 금액",
                    "type": "integer"
                },
                "productID": {
                    "type": "string"
                },
                "restCount": {
                    "description": "남은 수량",
                    "type": "integer"
                },
                "startDate": {
                    "description": "예매 시작 날짜 epoch time",
                    "type": "integer"
                },
                "totalCount": {
                    "description": "총 수량",
                    "type": "integer"
                }
            }
        },
        "request.ReqWithdrawalUser": {
            "type": "object",
            "properties": {
                "userID": {
                    "type": "string"
                }
            }
        },
        "response.GetsProduct": {
            "type": "object",
            "properties": {
                "category": {
                    "description": "상품 카테고리",
                    "type": "string"
                },
                "description": {
                    "description": "상품 설명",
                    "type": "string"
                },
                "endDate": {
                    "description": "예매 종료 날짜 epoch time",
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "description": "이미지",
                    "type": "string"
                },
                "name": {
                    "description": "상품이름",
                    "type": "string"
                },
                "perAmount": {
                    "description": "상품 티켓 당 금액",
                    "type": "integer"
                },
                "restCount": {
                    "description": "남은 수량",
                    "type": "integer"
                },
                "startDate": {
                    "description": "예매 시작 날짜 epoch time",
                    "type": "integer"
                },
                "totalCount": {
                    "description": "총 수량",
                    "type": "integer"
                }
            }
        },
        "response.ResGetProduct": {
            "type": "object",
            "properties": {
                "category": {
                    "description": "상품 카테고리",
                    "type": "string"
                },
                "description": {
                    "description": "상품 설명",
                    "type": "string"
                },
                "endDate": {
                    "description": "예매 종료 날짜 epoch time",
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "description": "이미지",
                    "type": "string"
                },
                "name": {
                    "description": "상품이름",
                    "type": "string"
                },
                "perAmount": {
                    "description": "상품 티켓 당 금액",
                    "type": "integer"
                },
                "restCount": {
                    "description": "남은 수량",
                    "type": "integer"
                },
                "startDate": {
                    "description": "예매 시작 날짜 epoch time",
                    "type": "integer"
                },
                "totalCount": {
                    "description": "총 수량",
                    "type": "integer"
                }
            }
        },
        "response.ResGetsProduct": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.GetsProduct"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
