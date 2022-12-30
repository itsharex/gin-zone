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
        "/base/auth/wxlogin": {
            "get": {
                "description": "微信授权登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "微信授权登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户apiid",
                        "name": "code",
                        "in": "query",
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
        "/base/upload": {
            "post": {
                "description": "上传接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基础模块"
                ],
                "summary": "上传接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.FileInfo"
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/base.User"
                            }
                        }
                    }
                }
            },
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
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MainUser"
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
        "/base/users/info": {
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
        "/base/users/login": {
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
                        "name": "LoginForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResType"
                        }
                    }
                }
            }
        },
        "/base/users/password": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "修改密码接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "description": "需要上传的json",
                        "name": "LoginForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResType"
                        }
                    }
                }
            }
        },
        "/base/users/{id}": {
            "delete": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
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
            },
            "patch": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "修改用户接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "需要上传的json",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUser"
                        }
                    }
                }
            }
        },
        "/mobile/chat/friends": {
            "get": {
                "description": "好友列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "聊天模块"
                ],
                "summary": "好友列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User.ID",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResType"
                        }
                    }
                }
            }
        },
        "/mobile/chat/logs": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "聊天记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "聊天模块"
                ],
                "summary": "聊天记录",
                "parameters": [
                    {
                        "description": "需要上传的json",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChatLogQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ChatLog"
                            }
                        }
                    }
                }
            }
        },
        "/third/gushici": {
            "get": {
                "description": "今日古诗词",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "第三方模块"
                ],
                "summary": "今日古诗词",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.Third"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.FileInfo": {
            "type": "object",
            "properties": {
                "filename": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "base.Third": {
            "type": "object"
        },
        "base.User": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "avatar": {
                    "description": "用户头像",
                    "type": "string",
                    "example": ":https://******.com/aa.png"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "123456789@qq.com"
                },
                "gender": {
                    "description": "性别",
                    "type": "integer"
                },
                "id": {
                    "description": "自增id",
                    "type": "integer"
                },
                "intro": {
                    "description": "个人介绍",
                    "type": "string",
                    "example": "个人介绍"
                },
                "isAdmin": {
                    "description": "是否管理员",
                    "type": "boolean",
                    "default": false
                },
                "isLock": {
                    "description": "是否已锁",
                    "type": "boolean",
                    "default": false
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string",
                    "example": "酱"
                },
                "password": {
                    "description": "密码 - 不会json化",
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 6,
                    "example": "123456"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userId": {
                    "description": "用户唯一id",
                    "type": "string"
                },
                "userName": {
                    "description": "用户名",
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 4,
                    "example": "test"
                }
            }
        },
        "model.ChangePassword": {
            "type": "object",
            "required": [
                "newPassword",
                "password",
                "userName"
            ],
            "properties": {
                "newPassword": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 6,
                    "example": "123456"
                },
                "password": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 6,
                    "example": "123456"
                },
                "userName": {
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 4,
                    "example": "test"
                }
            }
        },
        "model.ChatLog": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "groupId": {
                    "type": "integer"
                },
                "id": {
                    "description": "自增id",
                    "type": "integer"
                },
                "logType": {
                    "type": "integer"
                },
                "msgType": {
                    "type": "integer"
                },
                "receiverId": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "senderId": {
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "model.ChatLogQuery": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": ""
                },
                "groupId": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "pageSize": {
                    "type": "integer",
                    "example": 20
                },
                "receiverId": {
                    "type": "string"
                },
                "senderId": {
                    "type": "string"
                }
            }
        },
        "model.LoginForm": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 6,
                    "example": "123456"
                },
                "userName": {
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 4,
                    "example": "test"
                }
            }
        },
        "model.MainUser": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "avatar": {
                    "description": "用户头像",
                    "type": "string",
                    "example": ":https://******.com/aa.png"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "123456789@qq.com"
                },
                "gender": {
                    "description": "性别",
                    "type": "integer"
                },
                "intro": {
                    "description": "个人介绍",
                    "type": "string",
                    "example": "个人介绍"
                },
                "isAdmin": {
                    "description": "是否管理员",
                    "type": "boolean",
                    "default": false
                },
                "isLock": {
                    "description": "是否已锁",
                    "type": "boolean",
                    "default": false
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string",
                    "example": "酱"
                },
                "password": {
                    "description": "密码 - 不会json化",
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 6,
                    "example": "123456"
                },
                "userId": {
                    "description": "用户唯一id",
                    "type": "string"
                },
                "userName": {
                    "description": "用户名",
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 4,
                    "example": "test"
                }
            }
        },
        "model.UpdateUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "用户头像",
                    "type": "string",
                    "example": ":https://******.com/aa.png"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "123456789@qq.com"
                },
                "gender": {
                    "description": "性别",
                    "type": "integer"
                },
                "intro": {
                    "description": "个人介绍",
                    "type": "string",
                    "example": "个人介绍"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string",
                    "example": "酱"
                }
            }
        },
        "response.ResType": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "default": 0
                },
                "data": {},
                "msg": {
                    "type": "string",
                    "default": "操作成功"
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
