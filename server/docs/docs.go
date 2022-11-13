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
        "/base/delete/{id}": {
            "delete": {
                "description": "删除用户接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.User"
                        }
                    }
                }
            }
        },
        "/base/login": {
            "post": {
                "description": "用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "description": "需要上传的json",
                        "name": "user",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/base.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.User"
                        }
                    }
                }
            }
        },
        "/base/register": {
            "post": {
                "description": "用户注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "注册接口",
                "parameters": [
                    {
                        "description": "需要上传的json",
                        "name": "user",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/base.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.User"
                        }
                    }
                }
            }
        },
        "/base/users": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "用户列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "format": "email",
                        "description": "username search by q",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.User"
                        }
                    }
                }
            }
        },
        "/base/users/{id}": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "用户信息接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.User": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isAdmin": {
                    "description": "是否管理员",
                    "type": "boolean"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authorization": {
            "description": "jwt token 鉴权",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9600",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Gin-Zone-Api",
	Description:      "zone服务端API服务",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
