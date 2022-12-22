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
        "termsOfService": "https://tos.iexceed.dev",
        "contact": {
            "name": "Iexceed technology solutions",
            "url": "https://www.i-exceed.com/contact-us/",
            "email": "rahul.r@i-exceed.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "get": {
                "description": "Responds with the list of all accounts as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Get Accounts based on pageId and size",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Provide the pageId from where the records needs to be returned",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "provide the size of the page",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ListAccountRequest"
                        }
                    },
                    "400": {
                        "description": "Bad/Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Takes a account JSON and store in DB. Return saved JSON.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Create a new account",
                "parameters": [
                    {
                        "description": "Account JSON",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/CreateAccountRequest"
                        }
                    },
                    "400": {
                        "description": "Bad/Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "description": "Returns the account whose id value matches the isbn.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Get single account by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "search account by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Account"
                        }
                    },
                    "400": {
                        "description": "Bad/Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateAccountRequest": {
            "type": "object",
            "required": [
                "currency",
                "owner"
            ],
            "properties": {
                "currency": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "ListAccountRequest": {
            "type": "object",
            "required": [
                "pageID",
                "pageSize"
            ],
            "properties": {
                "pageID": {
                    "type": "integer",
                    "minimum": 1
                },
                "pageSize": {
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 5
                }
            }
        },
        "models.Account": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Fund transfer service",
	Description:      "A rest based service in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
