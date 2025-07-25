// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/login": {
            "post": {
                "description": "Login with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Authenticate user",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_HlufD_users-ms_internals_adapters_left_http_dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Authentication token",
                        "schema": {
                            "$ref": "#/definitions/github_com_HlufD_users-ms_internals_domain.Token"
                        }
                    },
                    "400": {
                        "description": "Invalid credentials",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/register": {
            "post": {
                "description": "Create a new user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Registration data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_HlufD_users-ms_internals_adapters_left_http_dto.RegisterUserDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created user",
                        "schema": {
                            "$ref": "#/definitions/github_com_HlufD_users-ms_internals_domain.User"
                        }
                    },
                    "400": {
                        "description": "Invalid request format",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/validate": {
            "post": {
                "description": "Validates a user based on token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Validate user token",
                "parameters": [
                    {
                        "description": "User Token Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_HlufD_users-ms_internals_adapters_left_http_dto.ValidateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user_id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_HlufD_users-ms_internals_adapters_left_http_dto.Login": {
            "description": "User login request payload",
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "@Example: \"P@ssw0rd123\"",
                    "type": "string",
                    "minLength": 5,
                    "example": "P@ssw0rd123"
                },
                "username": {
                    "description": "@Example: \"john_doe\"",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3,
                    "example": "john_doe"
                }
            }
        },
        "github_com_HlufD_users-ms_internals_adapters_left_http_dto.RegisterUserDto": {
            "description": "User registration request payload",
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "@Example: \"user@example.com\"",
                    "type": "string",
                    "format": "email",
                    "example": "user@example.com"
                },
                "password": {
                    "description": "@Example: \"P@ssw0rd123\"",
                    "type": "string",
                    "minLength": 6,
                    "example": "P@ssw0rd123"
                },
                "username": {
                    "description": "@Example: \"john_doe\"",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3,
                    "example": "john_doe"
                }
            }
        },
        "github_com_HlufD_users-ms_internals_adapters_left_http_dto.ValidateUser": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "github_com_HlufD_users-ms_internals_domain.Token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "github_com_HlufD_users-ms_internals_domain.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Users Microservice API",
	Description:      "API for user authentication and management",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
