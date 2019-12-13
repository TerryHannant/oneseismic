// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-11 16:22:52.2822329 +0100 CET m=+0.358001101

package docs

import (
	"bytes"
	"encoding/json"

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/manifest/{manifest_id}": {
            "get": {
                "description": "get manifest file",
                "produces": [
                    "application/octet-stream"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "File ID",
                        "name": "manifestID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Bytes"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "502": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.APIError"
                        }
                    }
                }
            },
            "post": {
                "description": "post manifest",
                "produces": [
                    "application/octet-stream"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "File ID",
                        "name": "manifestID",
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
                    },
                    "502": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/stitch/{manifest_id}/dim/{dim}": {
            "get": {
                "description": "post surface query to stitch",
                "produces": [
                    "application/octet-stream"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Bytes"
                        }
                    },
                    "400": {
                        "description": "Surface id not found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.APIError"
                        }
                    }
                }
            }
        },
        "/stitch/{manifest_id}/{surface_id}": {
            "get": {
                "description": "post surface query to stitch",
                "produces": [
                    "application/octet-stream"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Bytes"
                        }
                    },
                    "400": {
                        "description": "Surface id not found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.APIError"
                        }
                    }
                }
            }
        },
        "/surface/": {
            "get": {
                "description": "get list of available surfaces",
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.APIError"
                        }
                    }
                }
            }
        },
        "/surface/{surface_id}": {
            "get": {
                "description": "get surface file",
                "produces": [
                    "application/octet-stream"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "File ID",
                        "name": "surfaceID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Bytes"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.APIError"
                        }
                    },
                    "502": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.APIError"
                        }
                    }
                }
            },
            "post": {
                "description": "post surface file",
                "consumes": [
                    "application/octet-stream"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "File ID",
                        "name": "surfaceID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Bytes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.APIError": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                }
            }
        },
        "controller.Bytes": {
            "type": "array",
            "items": {}
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
var SwaggerInfo = swaggerInfo{Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
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
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
