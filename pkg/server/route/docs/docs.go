// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "contact": {
            "url": "https://nexclipper.io",
            "email": "jaehoon@nexclipper.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/client/regist": {
            "post": {
                "description": "Regist a Client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "parameters": [
                    {
                        "description": "Client의 정보",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqClient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/client/service": {
            "put": {
                "description": "Get a Service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "parameters": [
                    {
                        "description": "Service의 정보",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqClientGetService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.RespService"
                        }
                    }
                }
            }
        },
        "/server/catalogue": {
            "get": {
                "description": "Get a Catalogues",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Catalogues"
                        }
                    }
                }
            }
        },
        "/server/cluster": {
            "post": {
                "description": "Create a Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "parameters": [
                    {
                        "description": "Cluster의 정보",
                        "name": "cluster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqCluster"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/server/cluster/{id}": {
            "get": {
                "description": "Get a Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cluster의 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Cluster"
                        }
                    }
                }
            }
        },
        "/server/cluster/{id}/token": {
            "post": {
                "description": "Create a Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "cluster id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Token의 정보",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/server/service": {
            "post": {
                "description": "Create a Service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "parameters": [
                    {
                        "description": "Service의 정보",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Catalogue": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Catalogues": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Catalogue"
                    }
                }
            }
        },
        "model.Cluster": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "model.ReqClient": {
            "type": "object",
            "properties": {
                "cluster_id": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string"
                },
                "machine_id": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "model.ReqClientGetService": {
            "type": "object",
            "properties": {
                "cluster_id": {
                    "type": "integer"
                }
            }
        },
        "model.ReqCluster": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "model.ReqService": {
            "type": "object",
            "properties": {
                "cluster_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "step": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ReqStep"
                    }
                },
                "step_count": {
                    "type": "integer"
                }
            }
        },
        "model.ReqStep": {
            "type": "object",
            "properties": {
                "command": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parameter": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                }
            }
        },
        "model.ReqToken": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                }
            }
        },
        "model.RespService": {
            "type": "object",
            "properties": {
                "cluster_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "step": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RespStep"
                    }
                },
                "step_count": {
                    "type": "integer"
                }
            }
        },
        "model.RespStep": {
            "type": "object",
            "properties": {
                "command": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parameter": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
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
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "SUDORY",
	Description: "this is a sudory server.",
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
