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
        "/baskets": {
            "post": {
                "description": "return 201 if this could be created. Otherwise, it will return 500",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "basket"
                ],
                "summary": "Create a new basket.",
                "responses": {
                    "201": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "fail",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/baskets/{id}": {
            "get": {
                "description": "requires a basket ID example:\"0bfce8da-bdc9-11ec-b9f3-acde48001122\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "basket"
                ],
                "summary": "Show all products of basket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "requires a basket ID example:\"0bfce8da-bdc9-11ec-b9f3-acde48001122\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "basket"
                ],
                "summary": "remove a basket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/baskets/{id}/checkout": {
            "post": {
                "description": "requires a basket id, close of basket and will show details of order.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "basket"
                ],
                "summary": "close a basket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/baskets/{id}/products/{code}": {
            "post": {
                "description": "requires a basket id, and a product code. if product/code not exists then return \"product does not exist\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "basket"
                ],
                "summary": "add a new product to basket.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "CODE",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "requires a basket id, and a product code. if product/code not exists then return \"product does not exist\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "basket"
                ],
                "summary": "remove a product in the basket.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "CODE",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Item": {
            "type": "object",
            "properties": {
                "product": {
                    "$ref": "#/definitions/handler.Product"
                },
                "quantity": {
                    "type": "integer"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "handler.Product": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "handler.Response": {
            "type": "object",
            "properties": {
                "basket_id": {
                    "description": "basket id",
                    "type": "string"
                },
                "items": {
                    "description": "items",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.Item"
                    }
                },
                "total": {
                    "description": "total",
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "version(1.0)",
	Host:             "0.0.0.0:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API document title",
	Description:      "Description of specifications",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
