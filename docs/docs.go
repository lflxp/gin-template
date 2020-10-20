// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/admin/auth/add": {
            "post": {
                "description": "content 新增用户，不包括全权限",
                "tags": [
                    "Admin"
                ],
                "summary": "新增用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Auth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Auth"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/auth/all": {
            "get": {
                "description": "获取所有auth",
                "tags": [
                    "Admin"
                ],
                "summary": "获取所有auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/auth/del/{id}": {
            "delete": {
                "description": "删除Auth",
                "tags": [
                    "Admin"
                ],
                "summary": "删除Auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "更新的目标auth id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/auth/update/{id}": {
            "put": {
                "description": "content 修改用户名或密码",
                "tags": [
                    "Admin"
                ],
                "summary": "更新Auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "更新的目标auth id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Auth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Auth"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/auth/verify/{user}/{pwd}": {
            "get": {
                "description": "获取所有auth",
                "tags": [
                    "Admin"
                ],
                "summary": "获取所有auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "user",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "pwd",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Auth": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "creatorId": {
                    "type": "string"
                },
                "deleted": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "lastLoginIp": {
                    "type": "string"
                },
                "merchantCode": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "roleId": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "telephone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
