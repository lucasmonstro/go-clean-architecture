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
        "/todos": {
            "get": {
                "description": "Get a list of all todos",
                "tags": [
                    "todos"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.Todo"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new todo",
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "description": "input",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.CreateTodoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.Todo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.CreateTodoInput": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "schemas.Todo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "done": {
                    "type": "boolean"
                },
                "doneAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
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
