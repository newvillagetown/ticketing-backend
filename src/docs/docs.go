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
                            "$ref": "#/definitions/response.ResCallbackGoogleOAuth"
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
        "response.ResCallbackGoogleOAuth": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
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
