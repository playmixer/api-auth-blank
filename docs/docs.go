// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/login": {
            "post": {
                "description": "user authorization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "authorization",
                "parameters": [
                    {
                        "description": "User email",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.authorizeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.authorizeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    }
                }
            }
        },
        "/auth/otp": {
            "post": {
                "description": "send code to email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "authorization",
                "parameters": [
                    {
                        "description": "User email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "authorization",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.refreshRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.refreshResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "get user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "user data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.getUserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.responseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.authorizeRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "123456"
                },
                "email": {
                    "type": "string",
                    "example": "test@test.ru"
                }
            }
        },
        "api.authorizeResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.-5myAJwbMszwt7_iPciBQgICdujy20zKOZOUTXu9KyY"
                },
                "expires_in": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string",
                    "example": "213qwewq32q3q23qqrgrt67b54"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "api.getUserResponse": {
            "type": "object"
        },
        "api.loginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@test.ru"
                }
            }
        },
        "api.loginResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "api.refreshRequest": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string",
                    "example": "qweqwe231e2qeqae"
                }
            }
        },
        "api.refreshResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.-5myAJwbMszwt7_iPciBQgICdujy20zKOZOUTXu9KyY"
                },
                "expires_in": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string",
                    "example": "213qwewq32q3q23qqrgrt67b54"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "api.responseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean",
                    "example": false
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
